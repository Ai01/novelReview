package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var writeDB *gorm.DB
var readDB *gorm.DB

var writePing func(ctx context.Context) error

func legacyDSNFromEnv() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		return ""
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
}

func normalizeDSN(dsn string) string {
	return strings.TrimSpace(dsn)
}

func InitDB() error {
	writeDSN := normalizeDSN(os.Getenv("WRITE_DSN"))
	readDSN := normalizeDSN(os.Getenv("READ_DSN"))
	if writeDSN == "" {
		writeDSN = legacyDSNFromEnv()
	}
	if readDSN == "" {
		readDSN = writeDSN
	}

	var writeErr error
	writeDB, writePing, writeErr = connectWithRetries("write", writeDSN, 10)

	if readDSN != "" && readDSN != writeDSN {
		var readErr error
		readDB, _, readErr = connectWithRetries("read", readDSN, 5)
		if readErr != nil {
			log.Printf("Read database not available, fallback to write for reads: %v", readErr)
			readDB = nil
		}
	}

	return writeErr
}

func connectWithRetries(name string, dsn string, attempts int) (*gorm.DB, func(ctx context.Context) error, error) {
	if dsn == "" {
		return nil, nil, fmt.Errorf("%s dsn is empty", name)
	}

	var lastErr error
	for i := 0; i < attempts; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			lastErr = err
			log.Printf("Failed to open %s database (attempt %d): %v", name, i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}

		sqlDB, err := db.DB()
		if err != nil {
			lastErr = err
			log.Printf("Failed to get %s sql db (attempt %d): %v", name, i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		pingErr := sqlDB.PingContext(ctx)
		cancel()
		if pingErr != nil {
			lastErr = pingErr
			log.Printf("Failed to ping %s database (attempt %d): %v", name, i+1, pingErr)
			time.Sleep(2 * time.Second)
			continue
		}

		log.Printf("%s database connection established", name)
		return db, func(ctx context.Context) error {
			return sqlDB.PingContext(ctx)
		}, nil
	}

	return nil, nil, fmt.Errorf("final failure to connect to %s database after retries: %w", name, lastErr)
}

func DBForWrite() *gorm.DB {
	return writeDB
}

func DBForRead(strong bool) *gorm.DB {
	if strong {
		return writeDB
	}
	if readDB != nil {
		return readDB
	}
	return writeDB
}

func PingWrite(ctx context.Context) error {
	if writePing == nil {
		return fmt.Errorf("write database not initialized")
	}
	return writePing(ctx)
}

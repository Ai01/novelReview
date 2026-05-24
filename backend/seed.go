package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func seed() {
	var count int64
	DB.Model(&User{}).Where("username = ?", "admin").Count(&count)

	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		admin := User{
			Username: "admin",
			Email:    "admin@example.com",
			Password: string(hashedPassword),
			Avatar:   "https://api.dicebear.com/7.x/avataaars/svg?seed=admin",
		}
		if err := DB.Create(&admin).Error; err != nil {
			fmt.Printf("Failed to seed admin user: %v\n", err)
		} else {
			fmt.Println("Default admin user created: admin / admin123")
		}
	}
}

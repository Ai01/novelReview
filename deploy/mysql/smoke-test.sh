#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/../.."

MYSQL_ROOT_PASSWORD="${MYSQL_ROOT_PASSWORD:-rootpassword}"
MYSQL_DATABASE="${MYSQL_DATABASE:-novel_db}"

TOKEN="smoke_$(date +%s)"

docker compose -f deploy/docker-compose.yml exec -T mysql-master sh -lc "MYSQL_PWD='${MYSQL_ROOT_PASSWORD}' mysql -uroot \"${MYSQL_DATABASE}\" -e \"CREATE TABLE IF NOT EXISTS replication_smoke (id BIGINT AUTO_INCREMENT PRIMARY KEY, token VARCHAR(255) NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);\""
docker compose -f deploy/docker-compose.yml exec -T mysql-master sh -lc "MYSQL_PWD='${MYSQL_ROOT_PASSWORD}' mysql -uroot \"${MYSQL_DATABASE}\" -e \"INSERT INTO replication_smoke (token) VALUES ('${TOKEN}');\""

for i in $(seq 1 30); do
  if docker compose -f deploy/docker-compose.yml exec -T mysql-replica sh -lc "MYSQL_PWD='${MYSQL_ROOT_PASSWORD}' mysql -uroot \"${MYSQL_DATABASE}\" -e \"SELECT token FROM replication_smoke WHERE token='${TOKEN}' LIMIT 1;\"" | grep -q "${TOKEN}"; then
    exit 0
  fi
  sleep 1
done

exit 1

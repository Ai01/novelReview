#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/../.."

MYSQL_ROOT_PASSWORD="${MYSQL_ROOT_PASSWORD:-rootpassword}"
MYSQL_DATABASE="${MYSQL_DATABASE:-novel_db}"
MYSQL_REPLICATION_USER="${MYSQL_REPLICATION_USER:-repl}"
MYSQL_REPLICATION_PASSWORD="${MYSQL_REPLICATION_PASSWORD:-replpassword}"
MYSQL_BOOTSTRAP_TIMEOUT_SECS="${MYSQL_BOOTSTRAP_TIMEOUT_SECS:-480}"
DEFAULT_ADMIN_PASSWORD_BCRYPT="${DEFAULT_ADMIN_PASSWORD_BCRYPT:-\$2a\$10\$gVgIiefM5sApmshUBrILUO273tMIW/3UkvubLmYUzh30p2h7ROh.G}"

attempts_from_timeout() {
  local secs="$1"
  if [[ "${secs}" -lt 10 ]]; then
    echo 10
    return 0
  fi
  echo $(( secs / 2 ))
}

mysql_ok_with_password() {
  local service="$1"
  docker compose -f deploy/docker-compose.yml exec -T "${service}" sh -lc "MYSQL_PWD='${MYSQL_ROOT_PASSWORD}' mysql -uroot -e 'SELECT 1' >/dev/null 2>&1"
}

mysql_ok_without_password() {
  local service="$1"
  docker compose -f deploy/docker-compose.yml exec -T "${service}" sh -lc "MYSQL_PWD='' mysql -uroot -e 'SELECT 1' >/dev/null 2>&1"
}

ensure_mysql_root_password() {
  local service="$1"
  local max_attempts="$2"
  local saw_empty=0

  local start_ts
  start_ts="$(date +%s)"

  for _ in $(seq 1 "${max_attempts}"); do
    if [[ $(( _ % 10 )) -eq 0 ]]; then
      local now_ts
      now_ts="$(date +%s)"
      echo "[mysql] waiting ${service}... ($(( now_ts - start_ts ))s)"
    fi
    if mysql_ok_with_password "${service}"; then
      echo "[mysql] ${service} ok (root password)"
      return 0
    fi
    if mysql_ok_without_password "${service}"; then
      saw_empty=1
      echo "[mysql] ${service} root has empty password, setting root password..."
      docker compose -f deploy/docker-compose.yml exec -T "${service}" sh -lc "MYSQL_PWD='' mysql -uroot -e \"SET GLOBAL super_read_only=0; SET GLOBAL read_only=0; ALTER USER 'root'@'localhost' IDENTIFIED BY '${MYSQL_ROOT_PASSWORD}'; FLUSH PRIVILEGES;\""
      if mysql_ok_with_password "${service}"; then
        echo "[mysql] ${service} ok (root password set)"
        return 0
      fi
    fi
    sleep 2
  done

  echo "[mysql] ${service} not ready after ${MYSQL_BOOTSTRAP_TIMEOUT_SECS}s, tailing logs..."
  docker compose -f deploy/docker-compose.yml logs --tail 80 "${service}" || true

  if [[ "${saw_empty}" -eq 1 ]]; then
    echo "[mysql] ${service} root has empty password but failed to set expected password."
  fi

  echo "${service} root authentication failed."
  echo "Fix (reset dev data): sudo docker compose -f deploy/docker-compose.yml down -v --remove-orphans"
  return 1
}

echo "[mysql] ensuring mysql-master and mysql-replica root password..."
ensure_mysql_root_password mysql-master "$(attempts_from_timeout "${MYSQL_BOOTSTRAP_TIMEOUT_SECS}")"
ensure_mysql_root_password mysql-replica "$(attempts_from_timeout "${MYSQL_BOOTSTRAP_TIMEOUT_SECS}")"

echo "[mysql] creating replication user..."
docker compose -f deploy/docker-compose.yml exec -T mysql-master sh -lc "MYSQL_PWD='${MYSQL_ROOT_PASSWORD}' mysql -uroot -e \"CREATE USER IF NOT EXISTS '${MYSQL_REPLICATION_USER}'@'%' IDENTIFIED BY '${MYSQL_REPLICATION_PASSWORD}'; GRANT REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO '${MYSQL_REPLICATION_USER}'@'%'; FLUSH PRIVILEGES;\""

echo "[mysql] reading master status..."
master_status="$(docker compose -f deploy/docker-compose.yml exec -T mysql-master sh -lc "MYSQL_PWD='${MYSQL_ROOT_PASSWORD}' mysql -N -uroot -e \"SHOW MASTER STATUS;\"")"
master_status="$(echo "${master_status}" | tr -d '\r' | head -n 1)"

MASTER_LOG_FILE="$(echo "${master_status}" | awk '{print $1}')"
MASTER_LOG_POS="$(echo "${master_status}" | awk '{print $2}')"

if [[ -z "${MASTER_LOG_FILE}" || -z "${MASTER_LOG_POS}" ]]; then
  echo "Failed to read master binlog status. Ensure master binlog is enabled."
  exit 1
fi

echo "[mysql] configuring replica..."
docker compose -f deploy/docker-compose.yml exec -T mysql-replica sh -lc "MYSQL_PWD='${MYSQL_ROOT_PASSWORD}' mysql -uroot -e \"STOP REPLICA; RESET REPLICA ALL; CHANGE REPLICATION SOURCE TO SOURCE_HOST='mysql-master', SOURCE_USER='${MYSQL_REPLICATION_USER}', SOURCE_PASSWORD='${MYSQL_REPLICATION_PASSWORD}', SOURCE_LOG_FILE='${MASTER_LOG_FILE}', SOURCE_LOG_POS=${MASTER_LOG_POS}, GET_SOURCE_PUBLIC_KEY=1; START REPLICA;\""

replica_ready=0
for _ in $(seq 1 120); do
  if [[ $(( _ % 10 )) -eq 0 ]]; then
    echo "[mysql] waiting replica sync..."
  fi
  status="$(docker compose -f deploy/docker-compose.yml exec -T mysql-replica sh -lc "MYSQL_PWD='${MYSQL_ROOT_PASSWORD}' mysql -uroot -e \"SHOW REPLICA STATUS\\\\G\"" || true)"
  echo "${status}" | grep -qE "(Replica_IO_Running|Slave_IO_Running): Yes" || { sleep 1; continue; }
  echo "${status}" | grep -qE "(Replica_SQL_Running|Slave_SQL_Running): Yes" || { sleep 1; continue; }
  replica_ready=1
  break
done

if [[ "${replica_ready}" -ne 1 ]]; then
  echo "Replica did not become ready in time."
  exit 1
fi

echo "[mysql] applying schema and seeding data (on master)..."
docker compose -f deploy/docker-compose.yml exec -T mysql-master sh -lc "MYSQL_PWD='${MYSQL_ROOT_PASSWORD}' mysql -uroot -e \"CREATE DATABASE IF NOT EXISTS \\\`${MYSQL_DATABASE}\\\` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;\""

docker compose -f deploy/docker-compose.yml exec -T mysql-master sh -lc "MYSQL_PWD='${MYSQL_ROOT_PASSWORD}' mysql -uroot \"${MYSQL_DATABASE}\" < /bootstrap/schema.sql"

docker compose -f deploy/docker-compose.yml exec -T mysql-master env MYSQL_PWD="${MYSQL_ROOT_PASSWORD}" MYSQL_DATABASE="${MYSQL_DATABASE}" ADMIN_PWD_BCRYPT="${DEFAULT_ADMIN_PASSWORD_BCRYPT}" sh -lc 'mysql -uroot -e "INSERT INTO \`$MYSQL_DATABASE\`.users (username, email, password, avatar, created_at, updated_at) SELECT '\''admin'\'', '\''admin@example.com'\'', '\''$ADMIN_PWD_BCRYPT'\'', '\''https://api.dicebear.com/7.x/avataaars/svg?seed=admin'\'', NOW(), NOW() WHERE NOT EXISTS (SELECT 1 FROM \`$MYSQL_DATABASE\`.users WHERE username='\''admin'\'' LIMIT 1); UPDATE \`$MYSQL_DATABASE\`.users SET password='\''$ADMIN_PWD_BCRYPT'\'' WHERE username='\''admin'\'' AND password NOT LIKE '\''\\$2%'\'';"'

echo "[mysql] smoke test replication..."
bash ./deploy/mysql/smoke-test.sh

echo "[mysql] enforcing read-only on replica..."
docker compose -f deploy/docker-compose.yml exec -T mysql-replica sh -lc "MYSQL_PWD='${MYSQL_ROOT_PASSWORD}' mysql -uroot -e \"SET GLOBAL read_only=1; SET GLOBAL super_read_only=1;\""
exit $?

exit 1

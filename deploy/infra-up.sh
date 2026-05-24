#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

bash ./deploy/kill-ports.sh "${GATEWAY_PORT:-80}" "${PROMETHEUS_PORT:-19090}" "${GRAFANA_PORT:-3000}" "${LOKI_PORT:-3100}" "${MYSQL_MASTER_PORT:-33306}" "${MYSQL_REPLICA_PORT:-33307}"
docker compose -f deploy/docker-compose.yml up -d mysql-master mysql-replica loki promtail prometheus grafana nginx

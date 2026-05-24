#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

bash ./deploy/kill-ports.sh "${BACKEND_PORT:-18080}" "${FRONTEND_PORT:-5173}"
docker compose -f deploy/docker-compose.yml up -d --build backend
docker compose -f deploy/docker-compose.yml up -d --build frontend

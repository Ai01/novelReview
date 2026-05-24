#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

echo "[deploy] kill ports..."
bash ./deploy/kill-ports.sh
echo "[deploy] infra up..."
bash ./deploy/infra-up.sh
echo "[deploy] mysql bootstrap..."
bash ./deploy/bootstrap.sh
echo "[deploy] app up..."
bash ./deploy/app-up.sh
echo "[deploy] done."

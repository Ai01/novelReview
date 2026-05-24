#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

echo "[bootstrap] mysql bootstrap..."
bash ./deploy/mysql/bootstrap.sh
echo "[bootstrap] done."

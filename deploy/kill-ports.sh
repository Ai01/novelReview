#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

GATEWAY_PORT="${GATEWAY_PORT:-80}"
BACKEND_PORT="${BACKEND_PORT:-18080}"
FRONTEND_PORT="${FRONTEND_PORT:-5173}"
PROMETHEUS_PORT="${PROMETHEUS_PORT:-19090}"
GRAFANA_PORT="${GRAFANA_PORT:-3000}"
LOKI_PORT="${LOKI_PORT:-3100}"
MYSQL_MASTER_PORT="${MYSQL_MASTER_PORT:-33306}"
MYSQL_REPLICA_PORT="${MYSQL_REPLICA_PORT:-33307}"

default_ports=(
  "${GATEWAY_PORT}"
  "${BACKEND_PORT}"
  "${FRONTEND_PORT}"
  "${PROMETHEUS_PORT}"
  "${GRAFANA_PORT}"
  "${LOKI_PORT}"
  "${MYSQL_MASTER_PORT}"
  "${MYSQL_REPLICA_PORT}"
)

ports=()
if [[ "$#" -gt 0 ]]; then
  ports=("$@")
else
  ports=("${default_ports[@]}")
fi

stop_docker_publish_port() {
  local port="$1"
  if ! command -v docker >/dev/null 2>&1; then
    return 0
  fi

  local ids
  ids="$(docker ps -q --filter "publish=${port}" || true)"
  if [[ -n "${ids}" ]]; then
    docker stop ${ids} >/dev/null
  fi
}

pname_of_pid() {
  local pid="$1"
  ps -p "${pid}" -o comm= 2>/dev/null | tr -d '[:space:]'
}

pids_on_port() {
  local port="$1"
  if ! command -v ss >/dev/null 2>&1; then
    return 0
  fi

  ss -ltnp "sport = :${port}" 2>/dev/null | sed -n 's/.*pid=\([0-9]\+\).*/\1/p' | sort -u
}

kill_port() {
  local port="$1"

  stop_docker_publish_port "${port}"

  local pids
  pids="$(pids_on_port "${port}" || true)"
  if [[ -z "${pids}" ]]; then
    return 0
  fi

  for pid in ${pids}; do
    local pname
    pname="$(pname_of_pid "${pid}")"
    if [[ "${pname}" == "docker-proxy" || "${pname}" == "containerd-shim" || "${pname}" == "dockerd" ]]; then
      continue
    fi
    kill -TERM "${pid}" >/dev/null 2>&1 || true
  done

  sleep 1

  pids="$(pids_on_port "${port}" || true)"
  if [[ -z "${pids}" ]]; then
    return 0
  fi

  for pid in ${pids}; do
    local pname
    pname="$(pname_of_pid "${pid}")"
    if [[ "${pname}" == "docker-proxy" || "${pname}" == "containerd-shim" || "${pname}" == "dockerd" ]]; then
      continue
    fi
    kill -KILL "${pid}" >/dev/null 2>&1 || true
  done
}

for port in "${ports[@]}"; do
  kill_port "${port}"
done

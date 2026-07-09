#!/usr/bin/env bash
set -euo pipefail

# This wrapper exists because the official postgres image only drops privileges
# when the container command starts with `postgres`. Our previous `bash -c ...`
# caused postgres to be executed as root, which postgres refuses to do.

PGDATA="${PGDATA:-/var/lib/postgresql/data}"

if [[ -z "${POSTGRES_REPLICATION_USER:-}" ]]; then
  echo "ERROR: POSTGRES_REPLICATION_USER is required" >&2
  exit 1
fi
if [[ -z "${POSTGRES_REPLICATION_PASSWORD:-}" ]]; then
  echo "ERROR: POSTGRES_REPLICATION_PASSWORD is required" >&2
  exit 1
fi

# Ensure the data dir is writable by the postgres user (covers cases where an earlier run wrote as root).
if [[ "$(id -u)" == "0" ]]; then
  chown -R postgres:postgres "$PGDATA"
fi

export PGPASSWORD="${POSTGRES_REPLICATION_PASSWORD}"
until pg_isready -h "${POSTGRES_MASTER_HOST:-db-primary}" -p 5432 -U "${POSTGRES_REPLICATION_USER}"; do sleep 2; done

if [[ ! -s "${PGDATA}/PG_VERSION" ]]; then
  rm -rf "${PGDATA:?}/"*
  max_attempts=30
  attempt=1
  while ! su-exec postgres pg_basebackup -h db-primary -U "${POSTGRES_REPLICATION_USER}" -p 5432 -D "$PGDATA" -P -R -X stream -c fast; do
    if [[ "$attempt" -ge "$max_attempts" ]]; then
      echo "ERROR: pg_basebackup failed after ${max_attempts} attempts." >&2
      echo "HINT: Check primary has role '${POSTGRES_REPLICATION_USER}' with correct password." >&2
      echo "HINT: Re-init volumes to re-run init scripts: docker compose down -v && docker compose up -d" >&2
      exit 1
    fi
    attempt=$((attempt + 1))
    sleep 2
  done
  chown -R postgres:postgres "$PGDATA"
fi

exec /usr/local/bin/docker-entrypoint.sh "$@"


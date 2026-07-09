-- Single init script for the whole training stack.
-- Runs only on first init (empty docker volume). Re-init: `docker compose down -v`.

-- PgBouncer uses md5 auth with userlist hashes, so ensure role passwords are stored as MD5.
SET password_encryption = 'md5';

-- Minimal init: keep replica bootstrap working and keep passwords compatible with PgBouncer md5 userlist.
ALTER ROLE postgres WITH LOGIN PASSWORD 'admin';

-- 1) Replication user for db-replica pg_basebackup
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'replicator') THEN
    CREATE ROLE replicator WITH REPLICATION LOGIN PASSWORD 'admin';
  ELSE
    ALTER ROLE replicator WITH REPLICATION LOGIN PASSWORD 'admin';
  END IF;
END
$$;

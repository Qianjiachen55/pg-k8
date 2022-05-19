#!/bin/sh

set -e

echo "run db migration"

/app/tools/migrate -path /app/migration -database "$DB_Source" -verbose up

echo "start the app"
exec "$@"
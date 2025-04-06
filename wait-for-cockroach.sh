#!/bin/sh
host="$1"
port="$2"
shift 2

echo "⏳ Waiting for CockroachDB at $host:$port..."

while ! nc -z "$host" "$port"; do
  sleep 1
done

echo "✅ CockroachDB is up at $host:$port, launching app..."

echo "Running command: $@"

exec "$@"

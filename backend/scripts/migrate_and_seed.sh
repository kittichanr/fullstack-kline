#!/bin/bash

# Configuration
CONTAINER_NAME="mysql-db"
DB_NAME="kline"
DB_USER="root"
DB_PASSWORD="password"
SCHEMA_FILE="./db/migrations/schema.sql"
SEED_DIR="./db/seeds"  

# Check if the MySQL container is running
if !podman ps --format '{{.Names}}' | grep -q "$CONTAINER_NAME"; then
  echo "Error: MySQL container '$CONTAINER_NAME' is not running."
  exit 1
fi

# Apply schema migration
if [ -f "$SCHEMA_FILE" ]; then
  echo "Migrating schema..."
  podman exec -i $CONTAINER_NAME mysql -u$DB_USER -p$DB_PASSWORD $DB_NAME < $SCHEMA_FILE
else
  echo "Schema file '$SCHEMA_FILE' not found. Skipping migration."
fi

# Apply seed data from multiple files
if [ -d "$SEED_DIR" ]; then
  echo "Seeding mock data from directory: $SEED_DIR"
  for seed_file in $SEED_DIR/*.sql; do
    if [ -f "$seed_file" ]; then
      echo "Applying seed: $seed_file"
      podman exec -i $CONTAINER_NAME mysql -u$DB_USER -p$DB_PASSWORD $DB_NAME < "$seed_file"
    fi
  done
else
  echo "Seed directory '$SEED_DIR' not found. Skipping seeding."
fi

echo "Migration and seeding completed successfully!"

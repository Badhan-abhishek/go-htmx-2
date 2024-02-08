#!/bin/bash

set -e

# Replace 'your_database_name' with the actual name of your PostgreSQL database
DB_NAME="tasks"
DB_USER="postgres"
DB_PASSWORD="postgres"

# Drop the existing database
echo "Dropping database..."
psql -U $DB_USER -d postgres -c "DROP DATABASE IF EXISTS $DB_NAME;"

# Create a new database
echo "Creating database..."
psql -U $DB_USER -d postgres -c "CREATE DATABASE $DB_NAME;"

# Grant necessary privileges (optional)
# psql -U $DB_USER -d postgres -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;"

echo "Database $DB_NAME has been deleted and recreated successfully."
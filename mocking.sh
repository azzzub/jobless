#!/bin/bash

# Run this file by bash "./mocking.sh" command

echo "Droping all tables..."
go run main.go drop_all_tables

echo "Migrating databases..."
go run main.go migrate

echo "Creating user mockup..."
go run main.go user_mock

echo "Creating project mockup..."
go run main.go project_mock

echo "Creating bid mockup..."
go run main.go bid_mock

echo "FINISHED!"
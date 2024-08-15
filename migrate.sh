#!/bin/bash

# Ждем, пока PostgreSQL будет готов принимать соединения

migrate -path ./migrations -database postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable up

echo "Migrations completed."

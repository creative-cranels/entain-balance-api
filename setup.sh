#!/bin/bash

go install github.com/swaggo/swag/cmd/swag@v1.8.10
go install github.com/pressly/goose/v3/cmd/goose@latest

$GOPATH/bin/goose -dir "./db/migrations/" postgres "postgresql://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?search_path=$DB_SCHEMA" up

$GOPATH/bin/swag init -g cmd/main.go
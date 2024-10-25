source .env

goose -dir "./db/migrations/" postgres "postgresql://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?search_path=$DB_SCHEMA" $1

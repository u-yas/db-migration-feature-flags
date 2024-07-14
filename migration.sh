#!/bin/bash

db_service=$1

sql=$(cat ./alter_table_up.sql)

# docker compose exec経由でmigrationを実行する
if [[ $db_service = "mysql" ]]; then
  docker compose exec mysql mysql -u user -ppassword -D db -e "$sql"
elif [[ $db_service = "psql" ]]; then
  docker compose exec postgres psql -U user -d db -c "$sql"
else
  echo "Invalid argument"
  exit 1
fi

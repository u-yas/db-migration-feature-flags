#!/bin/bash
set -e pipefail
db_service=$1

if [[ $db_service = "mysql" ]]; then

  # 100回 migration_up.shとmigration_down.shを交互に実行する
  for i in {1..100}; do
    sql=$(cat ./alter_table_up.sql)
    docker compose exec mysql mysql -u user -ppassword -D db -e "$sql"
    sql=$(cat ./alter_table_down.sql)
    docker compose exec mysql mysql -u user -ppassword -D db -e "$sql"
  done
  exit 0
elif [[ $db_service = "psql" ]]; then
  # 100回 migration_up.shとmigration_down.shを交互に実行する
  for i in {1..100}; do
    sql=$(cat ./alter_table_up.sql)
    docker compose exec postgres psql -U user -d db -c "$sql"
    sql=$(cat ./alter_table_down.sql)
    docker compose exec postgres psql -U user -d db -c "$sql"
  done
  exit 0

else
  echo "Invalid argument"
  exit 1
fi

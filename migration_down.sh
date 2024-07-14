#!/bin/bash
db_service=$1

# alter_table.sqlからファイルの内容を読み取る
sql=$(cat ./alter_table_down.sql)

# docker compose exec経由でmigrationを実行する
if [[ $db_service = "mysql" ]]; then
  docker compose exec mysql mysql -u user -ppassword -D db -e "$sql"
elif [[ $db_service = "psql" ]]; then
  docker compose exec postgres psql -U user -d db -c "$sql"
else
  echo "Invalid argument"
  exit 1
fi

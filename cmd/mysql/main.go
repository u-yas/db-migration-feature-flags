package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "user"
	dbPass := "password"
	dbName := "db"

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 0.1秒間隔でgetUserを実行する
	for {
		err := getUser(db)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(10 * time.Millisecond)
	}

}

func getUser(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	var isEnabled bool
	err = tx.QueryRow("SELECT enabled FROM feature_flags WHERE name = 'users_name_to_display_name' FOR UPDATE").Scan(&isEnabled)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	if isEnabled {
		result, err := db.Query("SELECT id,display_name,age FROM users")
		if err != nil {
			return err
		}
		defer result.Close()

		for result.Next() {
			var id int
			var displayName string
			var age int
			err := result.Scan(&id, &displayName, &age)
			if err != nil {
				return err
			}

			fmt.Println("新しい")
		}
	} else {
		result, err := db.Query("SELECT id,name,age FROM users")
		if err != nil {
			return err
		}
		defer result.Close()
		for result.Next() {

			var id int
			var name string
			var age int
			err := result.Scan(&id, &name, &age)
			if err != nil {
				return err
			}
			fmt.Println("古い")
		}

	}
	return nil
}

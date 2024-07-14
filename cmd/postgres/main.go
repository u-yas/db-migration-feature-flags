package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	dbUser := "user"
	dbPass := "password"
	dbName := "db"

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPass))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
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
		return err
	}

	if isEnabled {
		result, err := tx.Query("SELECT id,display_name,age FROM users")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer result.Close()

		for result.Next() {
			var id int
			var displayName string
			var age int
			err := result.Scan(&id, &displayName, &age)
			if err != nil {
				tx.Rollback()
				return err
			}

			fmt.Println("新しい")
		}
	} else {
		result, err := tx.Query("SELECT id,name,age FROM users")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer result.Close()
		for result.Next() {

			var id int
			var name string
			var age int
			err := result.Scan(&id, &name, &age)
			if err != nil {
				tx.Rollback()
				return err
			}
			fmt.Println("古い")
		}

	}
	tx.Commit()
	return nil
}

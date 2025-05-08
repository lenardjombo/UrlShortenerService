package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func InitDB() (*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres",	"postgres://user_name:password@localhost/db_name?sslmode=disable")
	if err != nil {
		return nil, err
	} else {
		// Create model for our URL service
		stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS web_url (ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")

		if err != nil {
			log.Println(err)
			return nil, err
		}
		res, err := stmt.Exec()
		log.Println(res)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return db, nil
	}
}

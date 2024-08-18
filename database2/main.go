package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/chapter4")
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
}

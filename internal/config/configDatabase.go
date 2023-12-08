package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	DB, _ = sql.Open("mysql", "root:DevOps022019_@tcp(localhost:3306)/processing_db")

	err := DB.Ping()
	if err != nil {
		log.Fatal("Fatal connect to database!")
	}
}

func Disconnect() {
	DB.Close()
}

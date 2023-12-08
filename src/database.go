package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() *sql.DB {
	resultConnect, err := sql.Open("mysql", "root:DevOps022019_@tcp(127.0.0.1:3306)/testing_database")
	if err != nil {
		fmt.Println("Error connect to database!")
	}
	if err = resultConnect.Ping(); err != nil {
		if err != nil {
			fmt.Println("Error connect to database!")
		}
	}

	DB = resultConnect

	return resultConnect
}

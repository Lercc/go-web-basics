package config

import (
	"database/sql"
	"os"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))

	if err == nil  {
		fmt.Println("DB CONNECTION SUCCESS...")
	} else {
		fmt.Println("DB CONECCTION FAILED...")
	}

	return db
}
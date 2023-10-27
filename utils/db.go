package utils

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "gamal:Qu5dqVduSeh0@tcp(192.168.1.66:3306)/pionir")

	if err != nil {
		fmt.Println("Error ---", err)
		panic(err)
	}

	return db, err
}

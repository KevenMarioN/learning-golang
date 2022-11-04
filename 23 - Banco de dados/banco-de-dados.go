package main

import (
	"database/sql"
	"fmt"
	"log"

	dbConfig "banco-dados/dbconfig"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func main() {

	fmt.Printf("Accessing %s ... \n", dbConfig.DbName)

	db, err = sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection Open!")

	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println(rows)
}

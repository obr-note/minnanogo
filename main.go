package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("postgress", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

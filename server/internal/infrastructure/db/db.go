package db

import (
	"database/sql"
	"log"

	"github.com/dprio/currency-exchange/server/internal/infrastructure/db/dollarquote"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	DollarQuoteClient dollarquote.Client
}

func New() DB {
	db := startDB()

	return DB{
		DollarQuoteClient: dollarquote.NewClient(db),
	}
}

func startDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/currency_exchange?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

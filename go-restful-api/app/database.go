package app

import (
	"database/sql"
	"farhanhr/go-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go-migrate-restful-api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

//create migrate
// migrate create -ext sql -dir db/migration table_name

//run migrate
// migrate --database "mysql://root@tcp(localhost:3306)/go-migrate-restful-api" -path db/migrations up
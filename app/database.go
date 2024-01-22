package app

import (
	"RestAPIJWT/helper"
	"database/sql"
	"time"
)

func Database() *sql.DB {
	//conn := "postgres://postgres:Terserah123@localhost:5432/ginjwt?sslmode=disable"
	//mysql
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/test")
	helper.PanicError(err)

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB, DBNAHM, DBMASTER *sql.DB

func Conndbisoide() {
	dbConn, err := sql.Open("mysql", "adextux:adexganteng@tcp(localhost:3306)/pos_isoide")
	if err != nil {
		panic(err)
	}
	dbConn.SetMaxIdleConns(10)
	dbConn.SetMaxOpenConns(100)
	dbConn.SetConnMaxIdleTime(5 * time.Minute)
	dbConn.SetConnMaxLifetime(60 * time.Minute)
	DB = dbConn
}
func Conndbnahm() {
	dbConn, err := sql.Open("mysql", "adextux:adexganteng@tcp(localhost:3306)/pos_nahm")
	if err != nil {
		panic(err)
	}
	dbConn.SetMaxIdleConns(10)
	dbConn.SetMaxOpenConns(100)
	dbConn.SetConnMaxIdleTime(5 * time.Minute)
	dbConn.SetConnMaxLifetime(60 * time.Minute)
	DBNAHM = dbConn
}
func Connmaster() {
	dbConn, err := sql.Open("mysql", "adextux:adexganteng@tcp(localhost:3306)/pos_master")
	if err != nil {
		panic(err)
	}
	dbConn.SetMaxIdleConns(10)
	dbConn.SetMaxOpenConns(100)
	dbConn.SetConnMaxIdleTime(5 * time.Minute)
	dbConn.SetConnMaxLifetime(60 * time.Minute)
	DBMASTER = dbConn
}

package main

import (
	"database/sql"
	_ "gopkg.in/go-sql-driver/mysql.v1"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "user"
	dbPass := "user"
	dbName := "orderdb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}




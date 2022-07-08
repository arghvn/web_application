package config

// leraning details in future

import (
	"database/sql"
)

func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbName := "db6"
	dbUser := "root"
	dbPass := "123456"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return

}

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func getStringConnection(DBconfig *Database) string {
	stringConnection := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		DBconfig.USERNAME,
		DBconfig.PASSWORD,
		DBconfig.PROTOCOL,
		DBconfig.HOST,
		DBconfig.PORT,
		DBconfig.NAME)

	return stringConnection
}

func CreateConnection(DBconfig *Database) (*sql.DB, error) {
	db, err := sql.Open(DBconfig.DRIVER, getStringConnection(DBconfig))
	if err != nil {
		log.Errorw("Database open connection failed ", "Info", err.Error())
		return nil, err
	}
	return db, nil
}

func PingConnection(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		log.Errorw("Database ping failed", "Info", err.Error())
		return err
	}
	return nil
}

func InitConnection(DBconfig *Database) (*sql.DB, error) {
	db, err := CreateConnection(DBconfig)
	if err != nil {
		return nil, err
	}

	err = PingConnection(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

package connection

import (
	"database/sql"
	"fmt"

	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func getStringConnection(db *config.DBConfig) string {
	stringConnection := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true",
		db.USERNAME,
		db.PASSWORD,
		db.PROTOCOL,
		db.HOST,
		db.PORT,
		db.NAME)

	return stringConnection
}

func createConnection(db *config.DBConfig) (*sql.DB, error) {
	conn, err := sql.Open(db.DRIVER, getStringConnection(db))
	if err != nil {
		common.Log.Errorw("Database connection failed", "info", err.Error())
		return nil, err
	}
	return conn, nil
}

func pingConnection(conn *sql.DB) error {
	err := conn.Ping()
	if err != nil {
		common.Log.Errorw("Database ping failed", "info", err.Error())
		return err
	}
	return nil
}

func InitConnection(c *config.Config) (*sql.DB, error) {
	conn, err := createConnection(c.DB)
	if err != nil {
		return nil, err
	}

	err = pingConnection(conn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

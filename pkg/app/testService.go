package app

import (
	"database/sql"

	"github.com/QOSQOs/UNIVeasier/pkg/db"
	"github.com/QOSQOs/UNIVeasier/pkg/model/test"
)

func ServiceGetTest(conn *sql.DB, name string) (test.Test, error) {
	model, err := db.GetTest(conn, name)
	if err != nil {
		return test.Test{}, err
	}
	return model, nil
}

func ServiceGetListTest(conn *sql.DB) ([]test.Test, error) {
	listModel, err := db.GetListTest(conn)
	if err != nil {
		return nil, err
	}
	return listModel, nil
}

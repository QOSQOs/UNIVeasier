package testDao

import (
	"database/sql"

	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/pkg/db/sqlComponents"
	"github.com/QOSQOs/UNIVeasier/pkg/model"
)

func AddTest(conn *sql.DB, m model.Test) error {
	_, err := conn.Exec("call AddTest(?, ?, ?)", m.Name, m.Age, m.Email)
	if err != nil {
		common.Log.Errorw("The query cannot be done", "info", err.Error())
		return err
	}
	return nil
}

func GetTest(conn *sql.DB, name string) (model.Test, error) {
	var m model.Test
	err := conn.QueryRow("call GetTest(?)", name).Scan(&m.ID, &m.Name, &m.Age, &m.Email)
	if err != nil {
		common.Log.Errorw("The query cannot be done", "info", err.Error())
		return model.Test{}, err
	}
	return m, nil
}

func GetListTest(conn *sql.DB) ([]model.Test, error) {
	sqlQuery, err := sqlComponents.NewQuery(conn, "tabletest", "SELECT")
	if err != nil {
		return nil, err
	}

	sqlExpression, err := sqlQuery.GetSQLQuery()
	if err != nil {
		return nil, err
	}

	res, err := conn.Query(sqlExpression)
	if err != nil {
		common.Log.Errorw("The query cannot be done", "info", err.Error())
		return nil, err
	}

	var listModel []model.Test
	for res.Next() {
		var m model.Test
		err = res.Scan(&m.ID, &m.Name, &m.Age, &m.Email)
		if err != nil {
			common.Log.Errorw("The record cannot be read", "info", err.Error())
			return nil, err
		}
		listModel = append(listModel, m)
	}

	return listModel, nil
}

package db

import (
	"database/sql"

	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/pkg/model/test"
)

func GetTest(conn *sql.DB, name string) (test.Test, error) {
	var model test.Test
	err := conn.QueryRow("call GetTest(?)", name).Scan(&model.ID, &model.Name, &model.Age, &model.Email)
	if err != nil {
		common.Log.Errorw("The query cannot be done", "info", err.Error())
		return test.Test{}, err
	}
	return model, nil
}

func GetListTest(conn *sql.DB) ([]test.Test, error) {
	// Execute query to return many registries
	res, err := conn.Query("call GetListTest()")
	if err != nil {
		common.Log.Errorw("The query cannot be done", "info", err.Error())
		return nil, err
	}

	var listModel []test.Test
	for res.Next() {
		var model test.Test
		err = res.Scan(&model.ID, &model.Name, &model.Age, &model.Email)
		if err != nil {
			common.Log.Errorw("The record cannot be read", "info", err.Error())
			return nil, err
		}
		listModel = append(listModel, model)
	}

	return listModel, nil
}

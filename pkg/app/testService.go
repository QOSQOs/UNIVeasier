package app

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/QOSQOs/UNIVeasier/internal/common"

	"github.com/QOSQOs/UNIVeasier/pkg/db"
	"github.com/QOSQOs/UNIVeasier/pkg/model/test"
)

func ServiceAddTest(conn *sql.DB, body io.Reader) error {
	buf, err := ioutil.ReadAll(body)
	if err != nil {
		common.Log.Errorw("Cannot read body", "info", err.Error())
	}

	var model test.Test
	err = json.Unmarshal(buf, &model)
	if err != nil {
		common.Log.Errorw("Body Unmarshal failed", "info", err.Error())
	}

	err = db.AddTest(conn, model)
	if err != nil {
		return err
	}
	return nil
}

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

package testService

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/QOSQOs/UNIVeasier/internal/common"

	"github.com/QOSQOs/UNIVeasier/pkg/db/testDao"
	"github.com/QOSQOs/UNIVeasier/pkg/model"
)

func AddTest(conn *sql.DB, body io.Reader) error {
	buf, err := ioutil.ReadAll(body)
	if err != nil {
		common.Log.Errorw("Cannot read body", "info", err.Error())
		return err
	}

	var m model.Test
	err = json.Unmarshal(buf, &m)
	if err != nil {
		common.Log.Errorw("Body Unmarshal failed", "info", err.Error())
		return err
	}

	err = testDao.AddTest(conn, m)
	if err != nil {
		return err
	}
	return nil
}

func GetTest(conn *sql.DB, name string) (model.Test, error) {
	m, err := testDao.GetTest(conn, name)
	if err != nil {
		return model.Test{}, err
	}
	return m, nil
}

func GetListTest(conn *sql.DB) ([]model.Test, error) {
	listModel, err := testDao.GetListTest(conn)
	if err != nil {
		return nil, err
	}
	return listModel, nil
}

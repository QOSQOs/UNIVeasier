package personService

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/QOSQOs/UNIVeasier/internal/common"

	"github.com/QOSQOs/UNIVeasier/pkg/db/personDao"
	"github.com/QOSQOs/UNIVeasier/pkg/model"
)

func AddPerson(conn *sql.DB, body io.Reader) error {
	buf, err := ioutil.ReadAll(body)
	if err != nil {
		common.Log.Errorw("Cannot read body", "info", err.Error())
		return err
	}

	var m model.Person
	err = json.Unmarshal(buf, &m)
	if err != nil {
		common.Log.Errorw("Body Unmarshal failed", "info", err.Error())
		return err
	}

	err = m.Validate()
	if err != nil {
		common.Log.Errorw("Invalid model", "info", err.Error(), "json: ", m)
		return err
	}

	err = personDao.AddPerson(conn, m)
	if err != nil {
		return err
	}
	return nil
}

func GetPersonById(conn *sql.DB, id int64) (model.Person, error) {
	m, err := personDao.GetPersonById(conn, id)
	if err != nil {
		return model.Person{}, err
	}
	return m, nil
}

func GetListPerson(conn *sql.DB) ([]model.Person, error) {
	listModel, err := personDao.GetListPerson(conn)
	if err != nil {
		return nil, err
	}
	return listModel, nil
}

package personService

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/pkg/db/personDao"
	"github.com/QOSQOs/UNIVeasier/pkg/model"

	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
)

func AddPerson(conn *sql.DB, body io.Reader) error {
	buf, err := ioutil.ReadAll(body)
	if err != nil {
		common.Log.Errorw("Cannot read body", "info", err.Error())
		return err
	}

	var person model.Person
	err = json.Unmarshal(buf, &person)
	if err != nil {
		common.Log.Errorw("Body Unmarshal failed", "info", err.Error())
		return err
	}

	err = person.Validate()
	if err != nil {
		common.Log.Errorw("Invalid person model", "info", err.Error(), "json: ", person)
		return err
	}

	err = personDao.AddPerson(conn, person)
	if err != nil {
		return err
	}
	return nil
}

func GetPersonById(conn *sql.DB, id int64) (model.Person, error) {
	person, err := personDao.GetPersonById(conn, id)
	if err != nil {
		return model.Person{}, err
	}
	return person, nil
}

func GetListPerson(conn *sql.DB) ([]model.Person, error) {
	personList, err := personDao.GetListPerson(conn)
	if err != nil {
		return nil, err
	}
	return personList, nil
}

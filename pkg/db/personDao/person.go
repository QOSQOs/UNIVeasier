package personDao

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/model"

	"database/sql"
)

func AddPerson(conn *sql.DB, m model.Person) error {
	_, err := conn.Exec("call InsertPerson(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		m.Id, m.Code, m.Gender, m.FirstName, m.LastName, m.Phone, m.Email, m.Avatar, m.Type, m.HomeCity,
		m.CurrentCity, m.Ethnic, m.Nationality, m.BirthDate, m.AdmissionYear, m.Period, m.IsVerified, m.DocVerifier,
		m.CreatedDate, m.LastModifiedDate)

	if err != nil {
		common.Log.Errorw(utils.FailedSQLQuery("InsertPerson"), "info", err.Error())
		return err
	}
	return nil
}

func GetPersonById(conn *sql.DB, id int64) (model.Person, error) {
	var m model.Person
	err := conn.QueryRow("call GetPersonById(?)", id).Scan(
		&m.Id, &m.Code, &m.Gender, &m.FirstName, &m.LastName, &m.Phone, &m.Email, &m.Avatar, &m.Type, &m.HomeCity,
		&m.CurrentCity, &m.Ethnic, &m.Nationality, &m.BirthDate, &m.AdmissionYear, &m.Period, &m.IsVerified, &m.DocVerifier,
		&m.CreatedDate, &m.LastModifiedDate)

	if err != nil {
		common.Log.Errorw(utils.FailedSQLQuery("GetPersonById"), "info", err.Error())
		return model.Person{}, err
	}
	return m, nil
}

func GetListPerson(conn *sql.DB) ([]model.Person, error) {
	res, err := conn.Query("call GetListPerson()")
	if err != nil {
		common.Log.Errorw(utils.FailedSQLQuery("GetListPerson"), "info", err.Error())
		return nil, err
	}

	var personList []model.Person
	for res.Next() {
		var m model.Person
		err = res.Scan(
			&m.Id, &m.Code, &m.Gender, &m.FirstName, &m.LastName, &m.Phone, &m.Email, &m.Avatar, &m.Type, &m.HomeCity,
			&m.CurrentCity, &m.Ethnic, &m.Nationality, &m.BirthDate, &m.AdmissionYear, &m.Period, &m.IsVerified, &m.DocVerifier,
			&m.CreatedDate, &m.LastModifiedDate)

		if err != nil {
			common.Log.Errorw("The record cannot be read", "info", err.Error())
			return nil, err
		}
		personList = append(personList, m)
	}

	return personList, nil
}

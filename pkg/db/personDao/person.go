package personDao

import (
	"database/sql"

	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/pkg/model"
)

func AddPerson(conn *sql.DB, m model.Person) error {
	_, err := conn.Exec("call InsertPerson(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		m.Id, m.Code, m.Gender, m.FirstName, m.LastName, m.Phone, m.Email, m.Avatar, m.Type, m.HomeCity,
		m.CurrentCity, m.Ethnic, m.Nationality, m.BirthDate, m.AdmissionYear, m.Period, m.IsVerified, m.DocVerifier,
		m.CreatedDate, m.LastModifiedDate)

	if err != nil {
		common.Log.Errorw("The query cannot be done", "info", err.Error())
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
		common.Log.Errorw("The query cannot be done", "info", err.Error())
		return model.Person{}, err
	}
	return m, nil
}

func GetListPerson(conn *sql.DB) ([]model.Person, error) {
	// Execute query to return many registries
	res, err := conn.Query("call GetListPerson()")
	if err != nil {
		common.Log.Errorw("The query cannot be done", "info", err.Error())
		return nil, err
	}

	var listModel []model.Person
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
		listModel = append(listModel, m)
	}

	return listModel, nil
}

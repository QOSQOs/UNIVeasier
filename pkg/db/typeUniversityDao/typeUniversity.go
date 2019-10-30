package typeUniversityDao

import (
	"database/sql"

	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/pkg/model"
)

func AddTypeUniversity(conn *sql.DB, m model.TypeUniversity) error {
	_, err := conn.Exec("call InsertTypeUniversity(?, ?, ?, ?, ?, ?, ?, ?)",
		m.Id, m.Name, m.Description, m.IsVerified, m.DocVerifier, m.CreatedDate, m.LastModifiedDate, m.CreatedBy)

	if err != nil {
		common.Log.Errorw("The query cannot be done", "info", err.Error())
		return err
	}
	return nil
}

func GetTypeUniversityById(conn *sql.DB, id int64) (model.TypeUniversity, error) {
	var m model.TypeUniversity
	err := conn.QueryRow("call GetTypeUniversityById(?)", id).Scan(
		&m.Id, &m.Name, &m.Description, &m.IsVerified, &m.DocVerifier, &m.CreatedDate, &m.LastModifiedDate, &m.CreatedBy)

	if err != nil {
		common.Log.Errorw("The query cannot be done", "info", err.Error())
		return model.TypeUniversity{}, err
	}
	return m, nil
}

func GetListTypeUniversity(conn *sql.DB) ([]model.TypeUniversity, error) {
	// Execute query to return many registries
	res, err := conn.Query("call GetListTypeUniversity()")
	if err != nil {
		common.Log.Errorw("The query cannot be done", "info", err.Error())
		return nil, err
	}

	var listModel []model.TypeUniversity
	for res.Next() {
		var m model.TypeUniversity
		err = res.Scan(
			&m.Id, &m.Name, &m.Description, &m.IsVerified, &m.DocVerifier, &m.CreatedDate, &m.LastModifiedDate, &m.CreatedBy)

		if err != nil {
			common.Log.Errorw("The record cannot be read", "info", err.Error())
			return nil, err
		}
		listModel = append(listModel, m)
	}

	return listModel, nil
}

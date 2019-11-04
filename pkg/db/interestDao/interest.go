package interestDao

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/model"

	"database/sql"
)

func AddInterest(conn *sql.DB, m model.Interest) error {
	_, err := conn.Exec("call InsertInterest(?, ?, ?, ?, ?, ?)",
		m.Id, m.Tag, m.IsSkill, m.CreatedDate, m.LastModifiedDate, m.CreatedBy)

	if err != nil {
		common.Log.Errorw(utils.FailedSQLQuery("InsertInterest"), "info", err.Error())
		return err
	}
	return nil
}

func GetInterestById(conn *sql.DB, id int64) (model.Interest, error) {
	var m model.Interest
	err := conn.QueryRow("call GetInterestById(?)", id).Scan(
		&m.Id, &m.Tag, &m.IsSkill, &m.CreatedDate, &m.LastModifiedDate, &m.CreatedBy)

	if err != nil {
		common.Log.Errorw(utils.FailedSQLQuery("GetInterestById"), "info", err.Error())
		return model.Interest{}, err
	}

	return m, nil
}

func GetListInterest(conn *sql.DB) ([]model.Interest, error) {
	res, err := conn.Query("call GetListInterest()")
	if err != nil {
		common.Log.Errorw(utils.FailedSQLQuery("GetListInterest"), "info", err.Error())
		return nil, err
	}

	var interestList []model.Interest
	for res.Next() {
		var m model.Interest
		err = res.Scan(
			&m.Id, &m.Tag, &m.IsSkill, &m.CreatedDate, &m.LastModifiedDate, &m.CreatedBy)

		if err != nil {
			common.Log.Errorw("The record cannot be read", "info", err.Error())
			return nil, err
		}
		interestList = append(interestList, m)
	}

	return interestList, nil
}

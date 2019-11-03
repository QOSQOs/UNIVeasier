package interestService

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/pkg/db/interestDao"
	"github.com/QOSQOs/UNIVeasier/pkg/model"
)

func AddInterest(conn *sql.DB, body io.Reader) error {
	buf, err := ioutil.ReadAll(body)
	if err != nil {
		common.Log.Errorw("Cannot read body", "info", err.Error())
		return err
	}

	var interest model.Interest
	err = json.Unmarshal(buf, &interest)
	if err != nil {
		common.Log.Errorw("Body Unmarshal failed", "info", err.Error())
		return err
	}

	err = interest.Validate()
	if err != nil {
		common.Log.Errorw("Invalid interest model", "info", err.Error(), "json: ", interest)
		return err
	}

	err = interestDao.AddInterest(conn, interest)
	if err != nil {
		return err
	}
	return nil
}

func GetInterestById(conn *sql.DB, id int64) (model.Interest, error) {
	interest, err := interestDao.GetInterestById(conn, id)
	if err != nil {
		return model.Interest{}, err
	}
	return interest, nil
}

func GetListInterest(conn *sql.DB) ([]model.Interest, error) {
	interestList, err := interestDao.GetListInterest(conn)
	if err != nil {
		return nil, err
	}
	return interestList, nil
}

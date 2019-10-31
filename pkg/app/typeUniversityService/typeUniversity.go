package typeUniversityService

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/pkg/db/typeUniversityDao"
	"github.com/QOSQOs/UNIVeasier/pkg/model"

	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
)

func AddTypeUniversity(conn *sql.DB, body io.Reader) error {
	buf, err := ioutil.ReadAll(body)
	if err != nil {
		common.Log.Errorw("Cannot read body", "info", err.Error())
		return err
	}

	var typeUniversity model.TypeUniversity
	err = json.Unmarshal(buf, &typeUniversity)
	if err != nil {
		common.Log.Errorw("Body Unmarshal failed", "info", err.Error())
		return err
	}

	err = typeUniversity.Validate()
	if err != nil {
		common.Log.Errorw("Invalid model", "info", err.Error(), "json: ", typeUniversity)
		return err
	}

	err = typeUniversityDao.AddTypeUniversity(conn, typeUniversity)
	if err != nil {
		return err
	}
	return nil
}

func GetTypeUniversityById(conn *sql.DB, id int64) (model.TypeUniversity, error) {
	typeUniversity, err := typeUniversityDao.GetTypeUniversityById(conn, id)
	if err != nil {
		return model.TypeUniversity{}, err
	}
	return typeUniversity, nil
}

func GetListTypeUniversity(conn *sql.DB) ([]model.TypeUniversity, error) {
	typeUniversityList, err := typeUniversityDao.GetListTypeUniversity(conn)
	if err != nil {
		return nil, err
	}
	return typeUniversityList, nil
}

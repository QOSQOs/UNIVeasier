package albumService

import (
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/pkg/db/albumDao"
	"github.com/QOSQOs/UNIVeasier/pkg/model"
)

func AddAlbum(conn *sql.DB, body io.Reader) error {
	buf, err := ioutil.ReadAll(body)
	if err != nil {
		common.Log.Errorw("Cannot read body", "info", err.Error())
		return err
	}

	var albumModel model.Album
	err = json.Unmarshal(buf, &albumModel)
	if err != nil {
		common.Log.Errorw("Body Unmarshal failed", "info", err.Error())
		return err
	}

	err = albumModel.Validate()
	if err != nil {
		common.Log.Errorw("Invalid album model", "info", err.Error(), "json: ", albumModel)
		return err
	}

	err = albumDao.AddAlbum(conn, albumModel)
	if err != nil {
		return err
	}
	return nil
}

func GetAlbumById(conn *sql.DB, id int64) (model.Album, error) {
	albumModel, err := albumDao.GetAlbumById(conn, id)
	if err != nil {
		return model.Album{}, err
	}
	return albumModel, nil
}

func GetListAlbum(conn *sql.DB) ([]model.Album, error) {
	AlbumList, err := albumDao.GetListAlbum(conn)
	if err != nil {
		return nil, err
	}
	return AlbumList, nil
}

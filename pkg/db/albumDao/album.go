package albumDao

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/utils"
	"github.com/QOSQOs/UNIVeasier/pkg/model"

	"database/sql"
)

func AddAlbum(conn *sql.DB, albumModel model.Album) error {
	_, err := conn.Exec("call InsertAlbum(?, ?, ?, ?, ?, ?)",
		albumModel.Id, albumModel.Name, albumModel.Description, albumModel.CreatedDate, albumModel.LastModifiedDate, albumModel.CreatedBy)

	if err != nil {
		common.Log.Errorw(utils.FailedSQLQuery("InsertAlbum"), "info", err.Error())
		return err
	}
	return nil
}

func GetAlbumById(conn *sql.DB, id int64) (model.Album, error) {
	var albumModel model.Album
	err := conn.QueryRow("call GetAlbumById(?)", id).Scan(
		&albumModel.Id, &albumModel.Name, &albumModel.Description, &albumModel.CreatedDate, &albumModel.LastModifiedDate, &albumModel.CreatedBy)

	if err != nil {
		common.Log.Errorw(utils.FailedSQLQuery("GetAlbumById"), "info", err.Error())
		return model.Album{}, err
	}
	return albumModel, nil
}

func GetListAlbum(conn *sql.DB) ([]model.Album, error) {
	res, err := conn.Query("call GetListAlbum()")

	if err != nil {
		common.Log.Errorw(utils.FailedSQLQuery("GetListAlbum"), "info", err.Error())
		return nil, err
	}

	var albumList []model.Album
	for res.Next() {
		var albumModel model.Album
		err = res.Scan(
			&albumModel.Id, &albumModel.Name, &albumModel.Description, &albumModel.CreatedDate, &albumModel.LastModifiedDate, &albumModel.CreatedBy)

		if err != nil {
			common.Log.Errorw("The record cannot be read", "info", err.Error())
			return nil, err
		}
		albumList = append(albumList, albumModel)
	}

	return albumList, nil
}

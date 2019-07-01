package services

import (
	"bookTrade-backend/app"
	"bookTrade-backend/dao"
	"bookTrade-backend/models"
	"github.com/google/uuid"
	"io/ioutil"
	"mime/multipart"
)

func StoreImage(image *multipart.FileHeader) (string, error) {
	fileID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	if f, err := image.Open(); err != nil {
		return "", err
	} else {
		if data, err := ioutil.ReadAll(f); err != nil {
			return "", err
		} else if err := app.App.StorageManager.Store(fileID.String(), data); err != nil {
			return "", err
		}
	}
	// store map to database
	err = dao.CreateFileRef(&models.FileRef{
		FileKey:  fileID.String(),
		FileName: image.Filename,
	})
	// TODO: delete
	if err != nil {
		return "", err
	}

	return fileID.String(), nil
}

func FetchImage(key string) ([]byte, error) {
	return app.App.StorageManager.Fetch(key)
}

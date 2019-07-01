package dao

import (
	"bookTrade-backend/models"
	"github.com/jinzhu/gorm"
)

func CreateCatalog(catalog *models.Catalog) error {
	return db.Create(catalog).Error
}

func CreatCatalogIfNotExists(catalog *models.Catalog) error {
	c := models.Catalog{}
	if err := db.Where("name = ?", catalog.Name).First(&c).Error; err != nil {
		if gorm.IsRecordNotFoundError(err){
			return CreateCatalog(catalog)
		}else {
			return err
		}
	}
	return nil
	// found, do nothing
}

func GetAllCatalogs() ([]models.Catalog, error) {
	var catalogs []models.Catalog
	err := db.Find(&catalogs).Error
	return catalogs, err
}

func GetCatalogByID(catalogID int) (*models.Catalog, error) {
	var catalog models.Catalog
	err := db.Where("id = ?", catalogID).First(&catalog).Error
	return &catalog, err
}
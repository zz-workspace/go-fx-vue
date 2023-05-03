package repository

import (
	"fast-api.io/models"
	"gorm.io/gorm"
)

type ApiRepository struct {
	Db *gorm.DB
}

func InitApiRepository(Db *gorm.DB) *ApiRepository {
	repository := &ApiRepository{
		Db: Db,
	}
	return repository
}

func (c ApiRepository) CreateApi(api *models.API) *models.API {
	c.Db.Create(&api)
	return api
}

func (c ApiRepository) ListAPI(workspaceID uint64) []*models.API {
	var apis []*models.API
	c.Db.Where(&models.API{
		WorkspaceID: workspaceID,
	}).Find(&apis)
	return apis
}

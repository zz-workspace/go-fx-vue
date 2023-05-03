package repository

import (
	"fast-api.io/models"
	"gorm.io/gorm"
)

type EndpointRepository struct {
	Db *gorm.DB
}

func InitEndpointRepository(db *gorm.DB) *EndpointRepository {
	repository := &EndpointRepository{
		Db: db,
	}
	return repository
}

func (c EndpointRepository) CreateEndpoint(endpoint *models.Endpoint) *models.Endpoint {
	c.Db.Create(&endpoint)
	return endpoint
}

func (c EndpointRepository) FindBy(condition models.Endpoint) models.Endpoint {
	var endpoint models.Endpoint
	c.Db.Where(&condition).First(&endpoint)
	return endpoint
}

func (c EndpointRepository) FindEndpoint() []models.Endpoint {
	var endpoints []models.Endpoint
	c.Db.Find(&endpoints)

	return endpoints
}

func (c EndpointRepository) UpdateEndpoint(endpoint *models.Endpoint) *models.Endpoint {
	c.Db.Where(&models.Endpoint{
		ID: endpoint.ID,
	}).Updates(&endpoint)
	return endpoint
}

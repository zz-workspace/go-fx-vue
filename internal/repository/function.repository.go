package repository

import (
	"fast-api.io/helpers"
	"fast-api.io/models"
	"gorm.io/gorm"
)

type FunctionRepository struct {
	Db *gorm.DB
}

func InitFunctionRepository(db *gorm.DB) *FunctionRepository {
	repository := &FunctionRepository{
		Db: db,
	}
	return repository
}

func (c FunctionRepository) Create(function *models.Function) *models.Function {
	c.Db.Create(&function)
	return function
}

func (c FunctionRepository) FindById(id int) models.Function {
	var function models.Function

	c.Db.Where("id = ?", id).First(&function)

	return function
}

func (c FunctionRepository) Find() ([]models.Function, helpers.Paginator) {
	var functions []models.Function
	var paginator helpers.Paginator
	c.Db.Scopes(helpers.Paging(&paginator, &functions)).Find(&functions)
	return functions, paginator
}

func (c FunctionRepository) Update(function *models.Function) *models.Function {
	c.Db.Save(&function)
	return function
}

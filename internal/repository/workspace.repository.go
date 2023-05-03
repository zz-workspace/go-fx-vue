package repository

import (
	"fmt"

	"fast-api.io/helpers"
	"fast-api.io/models"
	"gorm.io/gorm"
)

type WorkspaceRepository struct {
	Db *gorm.DB
}

func InitWorkspaceRepository(Db *gorm.DB) *WorkspaceRepository {
	repository := &WorkspaceRepository{
		Db: Db,
	}
	return repository
}

func (c WorkspaceRepository) CreateWorkspace(workspace *models.Workspace) *models.Workspace {
	c.Db.Create(&workspace)
	return workspace
}

func (c WorkspaceRepository) DeleteManyWorkspaces(ids []uint64) []models.Workspace {
	fmt.Println(ids)
	var workspaces []models.Workspace
	c.Db.Delete(&workspaces, ids)
	return workspaces
}

func (c WorkspaceRepository) FindByName(name string) *models.Workspace {
	var workspace models.Workspace
	c.Db.Where(&models.Workspace{Name: name}).First(&workspace)
	return &workspace
}

func (c WorkspaceRepository) FindByID(ID uint64) *models.Workspace {
	var workspace models.Workspace
	c.Db.Where(&models.Workspace{ID: ID}).First(&workspace)
	return &workspace
}

func (c WorkspaceRepository) WorkspaceList() ([]models.Workspace, helpers.Paginator) {
	var workspaces []models.Workspace
	var paginator helpers.Paginator
	// c.Db.Scopes(helpers.Paging(&paginator, &tables)).Find(&tables)
	c.Db.Find(&workspaces)
	return workspaces, paginator
}

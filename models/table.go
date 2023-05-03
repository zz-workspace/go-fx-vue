package models

import (
	"gorm.io/datatypes"
)

type TableSchema struct {
	Name     string `json:"name"`
	DataType string `json:"data_type"`
}

type Table struct {
	ID          uint64         `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name        string         `json:"name"`
	Schema      datatypes.JSON `json:"schema" gorm:"type:jsonb"`
	WorkspaceID uint64         `json:"workspace_id"`
}

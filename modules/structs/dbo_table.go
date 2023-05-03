package structs

import "gorm.io/datatypes"

type Table struct {
	ID       uint64         `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name     string         `json:"name"`
	Schema   datatypes.JSON `json:"schema" gorm:"type:jsonb"`
	Contents datatypes.JSON `json:"contents" gorm:"type:jsonb;default:'[]'"`
}

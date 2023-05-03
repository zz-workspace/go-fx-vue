package models

type Workspace struct {
	ID          uint64  `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name        string  `json:"name"`
	DBServiceID string  `json:"db_service_id"`
	DBOrigin    string  `json:"db_origin"`
	Table       []Table `gorm:"foreignKey:WorkspaceID"`
}

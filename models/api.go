package models

type API struct {
	ID          uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name        string `json:"name"`
	WorkspaceID uint64 `json:"workspace_id"`
}

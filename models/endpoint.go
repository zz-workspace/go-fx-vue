package models

import "gorm.io/datatypes"

type method struct {
	Get    string
	Post   string
	Patch  string
	Put    string
	Delete string
}

var Method = &method{
	Get:    "GET",
	Post:   "POST",
	Patch:  "PATCH",
	Put:    "PUT",
	Delete: "DELETE",
}

type EndpointInput struct {
	Name     string `json:"name" validate:"required"`
	DataType string `json:"data_type" validate:"required"`
	Required bool   `json:"required" validate:"required"`
}

type EndpointFunction struct {
	Name    string         `json:"name" validate:"required"`
	Context datatypes.JSON `json:"context" gorm:"type:jsonb"`
	Input   datatypes.JSON `json:"input" gorm:"type:jsonb"`
	As      string         `json:"as" validate:"required"`
}

type EndpointResponse struct {
	Value string `json:"value" validate:"required"`
	As    string `json:"as" validate:"required"`
}

type Endpoint struct {
	ID          uint64         `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	WorkspaceID uint64         `json:"workspace_id"`
	ApiID       uint64         `json:"api_id"`
	Name        string         `json:"name"`
	Method      string         `json:"method"`
	Input       datatypes.JSON `json:"input" gorm:"type:jsonb"`
	Functions   datatypes.JSON `json:"functions" gorm:"type:jsonb"`
	Response    datatypes.JSON `json:"response" gorm:"type:jsonb"`
}

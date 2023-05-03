package models

import (
	"gorm.io/datatypes"
)

type FunctionName string

type FunctionTag string

const (
	Mvpuuid        FunctionName = "mvp::uuid"
	MvpDboGetby    FunctionName = "mvp::dbo_get_by"
	MvpDboGetAll   FunctionName = "mvp::dbo_getall"
	MvpDboUpdate   FunctionName = "mvp::dbo_update"
	MvpDboRawSql   FunctionName = "mvp::dbo_raw_sql"
	MvpTextAppend  FunctionName = "mvp::text_append"
	MvpNewVariable FunctionName = "mvp::new_variable"
)

const (
	Var   FunctionTag = "var"
	Const FunctionTag = "const"
)

type AppendContext struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
	Text  string `json:"text"`
}

type NewVariableContext struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type FunctionRun struct {
	As      string         `json:"as" validate:"required"`
	Name    FunctionName   `json:"name" validate:"required,oneof=mvp::uuid mvp::dbo_getby mvp::text_append"`
	Output  datatypes.JSON `json:"output"`
	Context datatypes.JSON `json:"context" validate:"required"`
}

type FunctionOutput struct {
	Name  string `json:"name" validate:"required,oneof=mvp::uuid mvp::dbo_getby mvp::text_append"`
	As    string `json:"as" validate:"required"`
	Value string `json:"value"`
}

type Function struct {
	ID      uint64         `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name    FunctionName   `json:"name" gorm:"not null"`
	Input   datatypes.JSON `json:"input" gorm:"type:jsonb;default:null"`
	Run     datatypes.JSON `json:"run" gorm:"type:jsonb;default:null"`
	Context datatypes.JSON `json:"context" gorm:"type:jsonb;default:null"`
	Output  datatypes.JSON `json:"output" gorm:"type:jsonb;default:null"`
}

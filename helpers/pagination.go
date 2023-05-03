package helpers

import (
	"gorm.io/gorm"
)

type Param struct {
	DB      *gorm.DB
	Page    int
	Limit   int
	OrderBy []string
	ShowSQL bool
}

type Paginator struct {
	TotalRecord int `json:"total_record"`
	TotalPage   int `json:"total_page"`
	// Records     interface{} `json:"records"`
	Offset   int `json:"offset"`
	Limit    int `json:"limit"`
	Page     int `json:"page"`
	PrevPage int `json:"prev_page"`
	NextPage int `json:"next_page"`
}

func Paging(paginator *Paginator, entity interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// var count int64
		// done := make(chan bool, 1)
		// go countRecords(db, entity, done, &count)

		return db.Offset(0).Limit(10)

	}
}

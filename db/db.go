package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"fast-api.io/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() *gorm.DB {
	dsn := fmt.Sprint(
		"host", "=", os.Getenv("MASTER_POSTGRES_HOST"),
		" ",
		"user", "=", os.Getenv("MASTER_POSTGRES_USER"),
		" ",
		"password", "=", os.Getenv("MASTER_POSTGRES_PASSWORD"),
		" ",
		"dbname", "=", os.Getenv("MASTER_POSTGRES_DB"),
		" ",
		"port", "=", os.Getenv("MASTER_POSTGRES_PORT"),
	)
	myLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: myLogger,
	})

	if err != nil {
		panic("Connect database failed")
	} else {
		db.AutoMigrate(&models.Function{})
		db.AutoMigrate(&models.Endpoint{})
		db.AutoMigrate(&models.Table{})
		db.AutoMigrate(&models.Workspace{})
		db.AutoMigrate(&models.API{})
	}

	return db
}

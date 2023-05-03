package main

// https://github.com/kecci/goscription/blob/master/app/cmd/http.go

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"fast-api.io/db"
	"fast-api.io/internal/controller"
	"fast-api.io/internal/repository"
	"fast-api.io/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(filepath.Join(os.Getenv("APP_PATH"), ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fx.New(
		fx.Options(
			fx.Provide(NewDB),
			fx.Provide(NewGin),
			repository.Module,
			controller.Module,
			service.Module,
		),
	).Run()
}

func NewDB(lc fx.Lifecycle) *gorm.DB {
	return db.NewDB()
}

// func NewRecordHandler(lc fx.Lifecycle, db *gorm.DB) *handlers.RecordHandler {

// }

func NewGin(lc fx.Lifecycle, db *gorm.DB) *gin.RouterGroup {
	r := gin.Default()
	r.Use(cors.Default())
	apiRoute := r.Group("api/v1")
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Print("Starting Gin server.")
			// apiRoutes := r.Group("api")
			// apiRoutes.GET("records", handlers.RecordHandler)
			go r.Run(":8888")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Print("Stopping Gin server")
			return nil
		},
	})
	return apiRoute
}

package main

import (
	"avito/internal/app"
	"avito/internal/models"
	"avito/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "host=localhost user=db_user password=pwd123 dbname=stats port=54320 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = models.InitModels(db)
	if err != nil {
		log.Fatal("too bad")
	}
	repo := repository.NewStatsRepository(db)
	application := app.NewApplication(repo)
	application.Start(":8080")
}

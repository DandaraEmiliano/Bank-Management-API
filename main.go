package main

import (
	"bank-management/api"
	"bank-management/models"
	"bank-management/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repository.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Account{})
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	r := gin.Default()
	api.SetupRoutes(r, db)

	r.Run(":8080")
}

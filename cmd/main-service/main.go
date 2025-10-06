package main

import (
	"fmt"
	"log"

	"github.com/waste3d/ADPP/internal/api/http/routers"
	v1 "github.com/waste3d/ADPP/internal/api/http/v1"
	"github.com/waste3d/ADPP/internal/config"
	"github.com/waste3d/ADPP/internal/domain"
	postgresStorage "github.com/waste3d/ADPP/internal/storage/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Println("Main Service is starting...")

	// 1 - configuration
	cfg := config.Load()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
	)

	// 2 - database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	err = db.AutoMigrate(&domain.Job{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// 3 - dependencies
	jobStorage := postgresStorage.New(db)
	handler := v1.NewHandler(jobStorage)

	router := routers.InitRouters(handler)
	log.Println("Router initialized.")

	log.Printf("Starting server on %s", cfg.Services.MainServicePort)
	if err := router.Run(":" + cfg.Services.MainServicePort); err != nil {
		log.Fatalf("FATAL: Failed to start server: %s", err)
	}
}

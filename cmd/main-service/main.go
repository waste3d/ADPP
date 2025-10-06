package main

import (
	"fmt"
	"log"

	"github.com/waste3d/ADPP/internal/api/http/routers"
	v1 "github.com/waste3d/ADPP/internal/api/http/v1"
	"github.com/waste3d/ADPP/internal/config"
	storage "github.com/waste3d/ADPP/internal/storage/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Main Service starting")

	cfg := config.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	jobStorage := storage.New(db)
	handler := v1.NewHandler(jobStorage)

	router := routers.InitRouters(jobStorage, handler)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Main Service started")
}

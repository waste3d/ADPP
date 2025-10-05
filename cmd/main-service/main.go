package main

import (
	"fmt"
	"log"

	"github.com/waste3d/ADPP/internal/config"
)

func main() {
	fmt.Println("Main Service starting")

	err, _ := config.Load()
	if err != nil {
		log.Fatalf("FATAL: Cannot read config: %s", err)
	}

	select {}
}

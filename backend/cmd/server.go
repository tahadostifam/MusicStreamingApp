package main

import (
	"fmt"
	"github.com/tahadostifam/MusicStreamingApp/config"
	"github.com/tahadostifam/MusicStreamingApp/database"
	"log"

	"github.com/tahadostifam/MusicStreamingApp/api"
)

func main() {
	cfg := config.Read()

	db := database.Connect(cfg.Database.DSN)

	apiErr := api.InitApi(cfg.App.HTTP.Host, cfg.App.HTTP.Port, db, cfg)
	if apiErr != nil {
		log.Fatalln("Failed to initialize api service!")
	}

	fmt.Printf("Server is listening on port %v\n", cfg.App.HTTP.Port)
}

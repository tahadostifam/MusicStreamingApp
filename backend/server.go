package main

import (
	"fmt"
	"log"

	"github.com/tahadostifam/MusicStreamingApp/api"
)

const PORT = 3001

func main() {
	apiErr := api.InitApi(PORT)
	if apiErr != nil {
		log.Fatalln("Failed to initialize api service!")
	}
	fmt.Printf("Server is listening on port %v\n", PORT)
}

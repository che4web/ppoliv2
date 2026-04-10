package main

import (
	"log"
	"os"

	"ppoliv2/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app, err := server.New()
	if err != nil {
		log.Fatalf("init app: %v", err)
	}

	if err = app.Run(":" + port); err != nil {
		log.Fatalf("run gin server: %v", err)
	}
}

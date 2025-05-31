package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"go-api-gateway/config"
	"go-api-gateway/router"
)

func main() {
	cfg := config.LoadConfig()

	app := fiber.New()

	router.SetupRoutes(app, cfg)

	log.Printf("Starting API Gateway on port %s...", cfg.Port)
	err := app.Listen(":" + cfg.Port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

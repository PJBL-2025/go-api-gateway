package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"go-api-gateway/config"
	"strings"
)

func SetupRoutes(app *fiber.App, cfg config.Config) {
	app.Use(logger.New())

	app.All("/go/*", func(c *fiber.Ctx) error {
		trimmedPath := strings.TrimPrefix(c.OriginalURL(), "/go")
		targetURL := cfg.GoBackendURL + trimmedPath
		return proxy.Do(c, targetURL)
	})

	app.All("/node/*", func(c *fiber.Ctx) error {
		trimmedPath := strings.TrimPrefix(c.OriginalURL(), "/node")
		targetURL := cfg.NodeBackendURL + trimmedPath
		return proxy.Do(c, targetURL)
	})
}

package routes

import (
	"github.com/ggorockee/working_with_microservice_in_go/cmd/api/handlers"
	swagger "github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Seoul",
	}))

	swaggerCfg := swagger.Config{
		BasePath: "/v1",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
	}

	app.Use(swagger.New(swaggerCfg))
	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		c.JSON(fiber.Map{
			"message": "🐣 v1",
		})
		return c.Next()
	})

	v1.Get("/healthz/ready", handlers.HealthCheck)
	return app
}

package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	Serve *fiber.App
}

func FiberNew(app *fiber.App) Server {
	return Server{
		Serve: app,
	}
}

func (app *Config) setupRoutes() {
	app.Server.Serve.Use(cors.New(
		cors.Config{
			AllowOrigins:     "*",
			AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
			AllowHeaders:     "Accept, Authorization, Content-Type, X-CSRF-Token",
			AllowCredentials: false,
			ExposeHeaders:    "Link",
			MaxAge:           300,
		},
	))

	app.Server.Serve.Post("/authenticate", app.Authenticate)
	app.Server.Serve.Get("/healthcheck", app.HealthCheck)
}

func (app *Config) Listen(aPort string) {
	if err := app.Server.Serve.Listen(fmt.Sprintf(":%s", aPort)); err != nil {
		panic(err)
	}
}

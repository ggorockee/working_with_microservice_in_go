package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App) {
	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     "*",
			AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
			AllowHeaders:     "Accept, Authorization, Content-Type, X-CSRF-Token",
			AllowCredentials: false,
			ExposeHeaders:    "Link",
			MaxAge:           300,
		},
	))

	app.Post("/authenticate", Authenticate)
}

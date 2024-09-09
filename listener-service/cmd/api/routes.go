package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (app *Config) Listen(aPort string) {
	if err := app.Server.Listen(fmt.Sprintf(":%s", aPort)); err != nil {
		panic(err)
	}
}

func (app *Config) setupRoutes() {
	app.Server.Use(cors.New(
		cors.Config{
			AllowOrigins:     "*",
			AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
			AllowHeaders:     "Accept, Authorization, Content-Type, X-CSRF-Token",
			AllowCredentials: false,
			ExposeHeaders:    "Link",
			MaxAge:           300,
		},
	))

}

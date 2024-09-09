package main

import (
	"github.com/ggorockee/working_with_microservice_in_go/authentication-service/data"
	"github.com/gofiber/fiber/v2"
)

const WEB_PORT = "80"

type Config struct {
	Server Server
}

func main() {
	data.ConnectDB()
	fiberApp := fiber.New()

	app := Config{
		Server: FiberNew(fiberApp),
	}

	app.setupRoutes()

	app.Listen(WEB_PORT)

}

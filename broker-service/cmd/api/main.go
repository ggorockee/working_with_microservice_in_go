package main

import (
	"github.com/gofiber/fiber/v2"
)

const (
	WEB_PORT = "80"
)

type Config struct {
	Server Server
}

func main() {
	fiberApp := fiber.New()
	app := Config{
		Server: FiberNew(fiberApp),
	}
	//app := fiber.New()
	//
	//setupRoutes(app)
	//
	//if err := app.Listen(fmt.Sprintf(":%s", webPort)); err != nil {
	//	panic(err)
	//}
	app.setupRoutes()
	app.Listen(WEB_PORT)
}

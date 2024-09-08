package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

const webPort = "80"

func main() {

	app := fiber.New()

	setupRoutes(app)

	if err := app.Listen(fmt.Sprintf(":%s", webPort)); err != nil {
		panic(err)
	}

}

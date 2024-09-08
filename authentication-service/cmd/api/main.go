package main

import (
	"fmt"
	"github.com/ggorockee/working_with_microservice_in_go/authentication-service/data"
	"github.com/gofiber/fiber/v2"
)

const webPort = "80"

func main() {

	//fiberConfig := fiber.Config{
	//	Prefork:      true,
	//	ServerHeader: "Fiber",
	//}

	app := fiber.New()

	//app = Config{
	//	srv:  fiber.New(fiberConfig),
	//	addr: fmt.Sprintf(":%s", webPort),
	//}

	setupRoutes(app)

	data.ConnectDB()

	if err := app.Listen(fmt.Sprintf(":%s", webPort)); err != nil {
		panic(err)
	}

}

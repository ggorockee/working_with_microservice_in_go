package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

const (
	WEBPORT string = "80"
)

func main() {
	engine := html.New("./cmd/web/templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	setupRoute(app)
	if err := app.Listen(fmt.Sprintf(":%s", WEBPORT)); err != nil {
		panic(err)
	}

}

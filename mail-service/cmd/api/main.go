package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	WEB_PORT = "80"
)

type Config struct {
	Server Server
	Mailer Mail
}

func main() {

	fiberApp := fiber.New()

	app := Config{
		Server: FiberNew(fiberApp),
		Mailer: createMail(),
	}

	// Routes
	app.setupRoutes()

	// Start server
	log.Println("Starting mail service on Port", WEB_PORT)
	app.Listen(WEB_PORT)
}

func createMail() Mail {
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	return Mail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        port,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Encryption:  os.Getenv("MAIL_ENCRYPTION"),
		FromAddress: os.Getenv("FROM_NAME"),
		FromName:    os.Getenv("FROM_ADDRESS"),
	}
}

package main

import (
	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math"
	"os"
	"time"
)

const (
	WEB_PORT = "80"
)

type Config struct {
	Server Server
	Rabbit *amqp.Connection
}

func main() {
	// Rabbit
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	fiberApp := fiber.New()
	app := Config{
		Server: FiberNew(fiberApp),
		Rabbit: rabbitConn,
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

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't continue until rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@broker-rabbit")
		if err != nil {
			log.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			log.Println("Connected to RabbitMQ!")
			connection = c
			break
		}

		if counts > 5 {
			log.Println("error:", err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
		continue
	}
	return connection, nil
}

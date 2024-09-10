package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"net"
	"net/rpc"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	WEB_PORT  = "80"
	RPC_PORT  = "5001"
	MONGO_URL = "mongodb://mongo:27017"
	gRPC_PORT = "50001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
	Server Server
}

func main() {
	// connect to mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Fiber
	fiberApp := fiber.New()

	app := Config{
		Models: data.New(client),
		Server: FiberNew(fiberApp),
	}

	// Register the RPC server
	err = rpc.Register(new(RPCServer))
	go app.rpcListen()

	// combine router
	app.setupRoutes()

	// start web server
	app.serve()
}

func (app *Config) rpcListen() error {
	log.Println("Starting RPC server on port ", RPC_PORT)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", RPC_PORT))
	if err != nil {
		return err
	}
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}

func (app *Config) serve() {
	if err := app.Server.Serve.Listen(fmt.Sprintf(":%s", WEB_PORT)); err != nil {
		panic(err)
	}
}

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(MONGO_URL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting: ", err)
		return nil, err
	}

	log.Println("Connected to mongodb")
	return c, nil
}

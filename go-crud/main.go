package main

import (
	"context"
	"os"

	"github.com/Wicho90/my-first-crud-go/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app := fiber.New()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/nest-pokemon"))
	if err != nil {
		panic(err)
	}

	defer disconnectDb(client)

	coll := client.Database("nest-pokemon").Collection("users")

	app.Use(cors.New())

	app.Static("/", "./public")

	controller.Routers(app, coll)

	app.Listen(":" + port)
}

func disconnectDb(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

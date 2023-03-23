package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routers(app *fiber.App, coll *mongo.Collection) {

	userController := NewUserController(coll)
	userGroup := app.Group("/users")

	userGroup.Get("", userController.GetAllUsers)
	userGroup.Post("", userController.CreateUser)

}

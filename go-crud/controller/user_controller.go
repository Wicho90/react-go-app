package controller

import (
	"context"

	"github.com/Wicho90/my-first-crud-go/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	*mongo.Collection
}

func NewUserController(coll *mongo.Collection) *UserController {
	return &UserController{coll}
}

func (u *UserController) GetAllUsers(c *fiber.Ctx) error {
	users := model.Users{}

	result, err := u.Collection.Find(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}

	for result.Next(context.TODO()) {
		user := model.NewUser()
		result.Decode(user)
		users = append(users, *user)
	}
	return c.JSON(&fiber.Map{
		"data": users,
	})
}

func (u *UserController) CreateUser(c *fiber.Ctx) error {
	user := model.NewUser()

	if err := c.BodyParser(user); err != nil {

	}

	result, err := u.Collection.InsertOne(context.TODO(), bson.D{{
		Key:   "name",
		Value: user.Name,
	}})

	if err != nil {
		panic(err)
	}

	return c.JSON(&fiber.Map{
		"data": result,
	})
}

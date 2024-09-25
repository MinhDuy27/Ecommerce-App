package handler

import (
	"fmt"
	rest "go-app/internal/api/rest"
	"net/http"

	"github.com/gofiber/fiber/v2"
)
type  UserHandler struct {}

func SetUpUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	handler := UserHandler{}

	app.Get("/users", handler.GetUser)
	app.Post("/users", handler.CreateUser)
}

func (u *UserHandler) GetUser(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "I'm a User",
	})
}
func (u *UserHandler) CreateUser(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User Created",
	})
}

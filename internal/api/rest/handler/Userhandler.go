package handler

import (
	_"fmt"
	rest "go-app/internal/api/rest"
	"go-app/internal/dto"
	"go-app/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)
type  UserHandler struct {
	usv service.UserService
}

func SetUpUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	
	usv := service.UserService{}
	handler := UserHandler{
		usv: usv,
	}
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
	// app.Get("/profiles", handler.GetProfiles)
	// app.Post("/profiles", handler.CreateProfiles)
	// app.Get("/carts", handler.GetCarts)
	// app.Post("/carts", handler.CreateCarts)
	// app.Get("/orders", handler.GetOrders)
	// app.Get("/orders/:id", handler.GetOrderByID)
	// app.Post("/become-seller", handler.BecomeSeller)
}

func (u *UserHandler) Register(c *fiber.Ctx) error {
	user := dto.SignUpdto{}
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "type valid input",
		})
	}
	value, error := u.usv.SignUp(user)
	if error != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error logging in",
		})
	}	
	
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": value,
	})
}
func (u *UserHandler) Login(c *fiber.Ctx) error {
	user := dto.Logindto{}
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "type valid input",
		})
	}
	message,error := u.usv.Login(user)

	if error != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "error creating user",
		})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": message,
	})
}
// func (u *UserHandler) GetProfiles(c *fiber.Ctx) error {
	
// }
// func (u *UserHandler) CreateProfiles(c *fiber.Ctx) error {
	
// }
// func (u *UserHandler) GetCarts(c *fiber.Ctx) error {
	
// }
// func (u *UserHandler) CreateCarts(c *fiber.Ctx) error {
	
// }
// func (u *UserHandler) GetOrders(c *fiber.Ctx) error {
	
// }
// func (u *UserHandler) GetOrderByID(c *fiber.Ctx) error {
	
// }
// func (u *UserHandler) BecomeSeller(c *fiber.Ctx) error {
	
// }
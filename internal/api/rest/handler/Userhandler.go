package handler

import (
	_ "fmt"
	"go-app/domain"
	rest "go-app/internal/api/rest"
	"go-app/internal/dto"
	"go-app/internal/repository"
	"go-app/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)
type  UserHandler struct {
	usv service.UserService
}

func SetUpUserRoutes(rh *rest.RestHandler) {
	app := rh.App
	Repo := repository.RepositoryImage(rh.Db) 
	usv := service.UserService{
		Repo: Repo,
	}
	handler := UserHandler{
		usv: usv,
	}
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
	app.Get("/user/id=:id",handler.GetProfilesbyID)
	app.Get("/user/email=:email",handler.GetProfilesbyEmail)
	app.Patch("/user/id=:id",handler.UpdateUser)
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
			"message": "error createing user",
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
	_,error := u.usv.Login(user)

	if error != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "error login",
		})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login success",
	})
}
func (u *UserHandler) GetProfilesbyID(c *fiber.Ctx) error {
	id := c.Params("id")
	result,error := u.usv.GetProfilesByID(id)
	if error != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "user not found",
	})	
	}	
	return   c.Status(http.StatusOK).JSON(result)
}
func (u *UserHandler) GetProfilesbyEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	result,error := u.usv.GetProfilesByEmail(email)
	if error != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "cannot find by your email",
	})	
	}	
	return   c.Status(http.StatusOK).JSON(result)
}
func (u *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	_,error := u.usv.GetProfilesByID(id)
	if error != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "there was an error",
	})}
	update := domain.User{}
    if err := c.BodyParser(&update); err != nil {
        return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
            "error": "invalid request body",
        })
    }
	err := u.usv.UpdateUser(id,update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
            "error": "cannot update user",
        })
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user updated",
	})
}
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
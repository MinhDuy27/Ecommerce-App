package handler

import (
	"encoding/json"
	_ "fmt"
	"github.com/MinhDuy27/go-app/domain"
	rest "github.com/MinhDuy27/go-app/internal/api/rest"
	"github.com/MinhDuy27/go-app/internal/dto"
	"github.com/MinhDuy27/go-app/internal/helper"
	"github.com/MinhDuy27/go-app/internal/repository"
	"github.com/MinhDuy27/go-app/internal/service"
	"log"
	"net/http"
	"strconv"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gofiber/fiber/v2"
)
type  UserHandler struct {
	usv service.UserService
	cache *helper.Client
}

func SetUpUserRoutes(rh *rest.RestHandler) {
	app := rh.App
	Repo := repository.RepositoryImage(rh.Db) 
	cached := rh.Cached
	usv := service.UserService{
		Repo: Repo,
		Auth: rh.Auth,
	}
	handler := UserHandler{
		usv: usv,
		cache: cached,
	}
	//public Endpoint
	pubRoute := app.Group("/users")

	pubRoute.Post("/register", handler.Register)
	pubRoute.Post("/login", handler.Login)


	//private Endpoint
	prvtRoute := pubRoute.Group("/",rh.Auth.Authorize)

	prvtRoute.Get("/profiles",rh.Cached.VerifyCache,handler.GetProfiles)
	prvtRoute.Patch("/update",handler.UpdateUser)
	prvtRoute.Post("/profiles", handler.CreateProfiles)
	prvtRoute.Patch("/sellers", handler.BecomeSeller)
	prvtRoute.Patch("/revoke-sellers", handler.RevokeSellerRole)
	prvtRoute.Get("/verifications", handler.GetVerificationCode)
	prvtRoute.Post("/verifications", handler.Verification)
	// app.Get("/carts", handler.GetCarts)
	// app.Post("/carts", handler.CreateCarts)
	// app.Get("/orders", handler.GetOrders)
	// app.Get("/orders/:id", handler.GetOrderByID)
}

func (u *UserHandler) Register(c *fiber.Ctx) error {
	user := dto.SignUpdto{}
	err := c.BodyParser(&user)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "type valid input",
		})
	}
	token, error := u.usv.SignUp(user)
	if error != nil {
		log.Println(error)
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error creating user",
		})
	}	
	
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register success",
		"token": token,
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
	token,error := u.usv.Login(user)

	if error != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "login failed",
		})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login success",
		"token": token,
	})
}
func (u *UserHandler) GetProfiles(c *fiber.Ctx) error {
	id := u.usv.Auth.GetUser(c).ID
	idStr := strconv.FormatUint(uint64(id), 10)
	result,err := u.usv.GetProfilesByID(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "user not found",
	})}
	resultByte, err := json.Marshal(result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "parse []byte error",
	})}
	cacheErr := u.cache.Client.Set(&memcache.Item{Key: idStr, Value: resultByte, Expiration: 100})
	if cacheErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "memcached error",
	})}
	return c.Status(http.StatusOK).JSON(result)
}
func (u *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := u.usv.Auth.GetUser(c).ID
	_,error := u.usv.GetProfilesByID(id)
	if error != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "there was an error when verify your id",
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
func (u *UserHandler) CreateProfiles(c *fiber.Ctx) error {
	id := u.usv.Auth.GetUser(c).ID
	_,error := u.usv.GetProfilesByID(id)
	if error != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "there was an error when verify your id",
	})}
	user := dto.CreateProfiledto{}
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "invalid request body",
		})
	}
	err1 := u.usv.CreateProfile(id,user)
	if err1 != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error creating profile",
		})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "profile created",
	})

}
func (u *UserHandler) BecomeSeller(c *fiber.Ctx) error {
	id := u.usv.Auth.GetUser(c).ID
	user,error := u.usv.GetProfilesByID(id)
	if error != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "there was an error when verify your id",
	})}
	if user.UserType == "seller" {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "you are already a seller",
		})
	}
	err := u.usv.BecomeSeller(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "there was an error when become seller",
		})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "you are now a seller",
	})
}
func (u *UserHandler) RevokeSellerRole(c *fiber.Ctx) error {
	id := u.usv.Auth.GetUser(c).ID
	user,error := u.usv.GetProfilesByID(id)
	if error != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "there was an error when verify your id",
	})}
	if user.UserType == "buyer" {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "you are already not a seller",
		})
	}
	err := u.usv.RevokeSeller(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "there was an error when revoke seller role",
		})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "revoke seller role success",
	})
}
func (u *UserHandler) GetVerificationCode(c *fiber.Ctx) error {
	id := u.usv.Auth.GetUser(c).ID
	user,error := u.usv.GetProfilesByID(id)
	if error != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "there was an error when verify your id",
	})}
	code,error := u.usv.GetVerificationCode(user)
	if error != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": error.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verification code",
		"code": code,
	})
}
func (u *UserHandler) Verification(c *fiber.Ctx) error {
	id := u.usv.Auth.GetUser(c).ID
	_,error := u.usv.GetProfilesByID(id)
	if error != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "there was an error when verify your id",
	})}

	input := dto.VerificationCode{}
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "invalid request body",
		})
	}
	err1 := u.usv.VerifyCode(id,input.Code)
	if err1 != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "error verify code",
		})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "code verified",
	})
}
// func (u *UserHandler) GetCarts(c *fiber.Ctx) error {
	
// }
// func (u *UserHandler) CreateCarts(c *fiber.Ctx) error {
	
// }
// func (u *UserHandler) GetOrders(c *fiber.Ctx) error {
	
// }
// func (u *UserHandler) GetOrderByID(c *fiber.Ctx) error {
	
// }
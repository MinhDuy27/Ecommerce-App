package  rest

import (
	"go-app/internal/helper"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestHandler struct {
	App *fiber.App
	Db *gorm.DB
	Auth helper.Auth
}
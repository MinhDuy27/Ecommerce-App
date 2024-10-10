package rest

import (
	"github.com/MinhDuy27/Ecommerce-App/internal/helper"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestHandler struct {
	App *fiber.App
	Db *gorm.DB
	Auth helper.Auth
	Cached *helper.Client
}
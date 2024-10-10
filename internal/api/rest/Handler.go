package rest

import (
	"github.com/MinhDuy27/go-app/internal/helper"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestHandler struct {
	App *fiber.App
	Db *gorm.DB
	Auth helper.Auth
	Cached *helper.Client
}
package api

import (
	"go-app/configs"
	rest "go-app/internal/api/rest"
	RestHandler "go-app/internal/api/rest/handler"
	"github.com/gofiber/fiber/v2"
)

func StartServer (config configs.AppConfig) {
	app:= fiber.New()  // create app
	rh := &rest.RestHandler{
		App : app,
	} // create handler
	SetUpRoute(rh) // use handler to set up routes
	app.Listen(config.ServerPort) // listen to port

}


// server function to set up routes
func SetUpRoute(rh *rest.RestHandler) {
	RestHandler.SetUpUserRoutes(rh)
}
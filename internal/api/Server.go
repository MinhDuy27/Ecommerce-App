package api

import (
	"go-app/configs"
	"go-app/domain"
	rest "go-app/internal/api/rest"
	RestHandler "go-app/internal/api/rest/handler"
	"log"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer (config configs.AppConfig) {
	app:= fiber.New()  // create app

	// connect to db
	db,err := gorm.Open(postgres.Open(config.Dsn),&gorm.Config{}) 
	if err != nil{
		log.Fatalf("Db connection failed %v",err)
	}
	log.Printf("DB connected")

	// create handler
	rh := &rest.RestHandler{
		App : app,
		Db : db,
	} 

	db.AutoMigrate(&domain.User{}) // migrate db
	
	SetUpRoute(rh) // use handler to set up routes
	
	app.Listen(config.ServerPort) // listen to port

}

// server function to set up routes
func SetUpRoute(rh *rest.RestHandler) {
	RestHandler.SetUpUserRoutes(rh)
}
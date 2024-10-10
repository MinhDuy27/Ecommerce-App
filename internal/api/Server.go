package api

import (
	"github.com/MinhDuy27/Ecommerce-App/configs"
	"github.com/MinhDuy27/Ecommerce-App/domain"
	rest "github.com/MinhDuy27/Ecommerce-App/internal/api/rest"
	RestHandler "github.com/MinhDuy27/Ecommerce-App/internal/api/rest/handler"
	"github.com/MinhDuy27/Ecommerce-App/internal/helper"
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
	Auth := helper.GetAuth(config.AppSecret)
	// create handler
	Cache,err := helper.NewMemcached()
	if err!= nil {
		log.Fatalf("error connecting to memcached %v",err)
	}
	
	rh := &rest.RestHandler{
		App : app,
		Db : db,
		Auth: Auth,
		Cached: Cache,
	} 

	db.AutoMigrate(&domain.User{}) // migrate db
	
	SetUpRoute(rh) // use handler to set up routes
	
	app.Listen(config.ServerPort) // listen to port

}

// server function to set up routes
func SetUpRoute(rh *rest.RestHandler) {
	RestHandler.SetUpUserRoutes(rh)
}
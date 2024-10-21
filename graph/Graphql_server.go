package graph

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MinhDuy27/Ecommerce-App/configs"
	"github.com/MinhDuy27/Ecommerce-App/domain"
	"github.com/MinhDuy27/Ecommerce-App/internal/helper"
	"github.com/MinhDuy27/Ecommerce-App/internal/repository"
	"github.com/MinhDuy27/Ecommerce-App/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func GraphServer(config configs.AppConfig) {
	db,err := gorm.Open(postgres.Open(config.Dsn),&gorm.Config{}) 
	if err != nil{
		log.Fatalf("Db connection failed %v",err)
	}
	db.AutoMigrate(
		&domain.User{},
		&domain.Product{},
	)
	
	port := os.Getenv(config.ServerPort)
	if port == "" {
		port = defaultPort
	}
	Repo := repository.RepositoryImage(db) 
	Auth := helper.GetAuth(config.AppSecret)
	usv := service.UserService{
		Repo: Repo,
		Auth: Auth,
	}
	ProductRepo := repository.GetProductImage(db)
	psv := service.ProductService{
		Rp: ProductRepo,
	}
	log.Printf("DB connected")
	
	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{
		Usv : usv,
		Psv : psv,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

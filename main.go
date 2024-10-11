package main

import (
	"github.com/MinhDuy27/Ecommerce-App/configs"
	"github.com/MinhDuy27/Ecommerce-App/graph"
	// "github.com/MinhDuy27/Ecommerce-App/internal/api"
	"log"
)

func main() {
	cfg, err := configs.SetUpEnv()
	if err != nil {
		log.Fatalf("there was an error setting up the environment %v\n", err)
	}
	// api.StartServer(cfg)
	graph.GraphServer(cfg)
}
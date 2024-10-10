package main

import (
	"github.com/MinhDuy27/go-app/configs"
	"github.com/MinhDuy27/go-app/internal/api"
	"log"
)

func main() {
	cfg, err := configs.SetUpEnv()
	if err != nil {
		log.Fatalf("there was an error setting up the environment %v\n", err)
	}
	api.StartServer(cfg)
}
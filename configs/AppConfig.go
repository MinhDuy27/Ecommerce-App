package configs

import (
	"errors"
	"os"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
}
func SetUpEnv()(cfg AppConfig , err error){
	// if os.Getenv("APP_ENV") == "DEV" {
	 	godotenv.Load()
	// }
	httpPort := os.Getenv("HTTP_PORT")

	if len(httpPort) == 0 {
		return AppConfig{}, errors.New("HTTP_PORT is not set")
	}
	return AppConfig{
		ServerPort: httpPort,
	}, nil
}
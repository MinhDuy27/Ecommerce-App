package configs

import (
	"errors"
	"os"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn string
	AppSecret string
}
func SetUpEnv()(cfg AppConfig , err error){
	
	godotenv.Load()
	HttpPort := os.Getenv("HTTP_PORT")
	if len(HttpPort) == 0 {
		return AppConfig{}, errors.New("HTTP_PORT is not set")
	}
	Dsn := os.Getenv("DSN")
	if len(Dsn) == 0 {
		return AppConfig{}, errors.New("DSN is not set")
	}
	appSecrect := os.Getenv("APP_SECRET")
	if len(appSecrect) == 0 {
		return AppConfig{}, errors.New("appSecrect is not set")
	}
	return AppConfig{
		ServerPort: HttpPort,
		Dsn : Dsn,
		AppSecret: appSecrect,
	}, nil
}
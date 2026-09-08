package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/projectdiscovery/gologger"
	"srebootcamp/internal/model"
)

func convert(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		gologger.Fatal().Msg(err.Error())
	}
	return i
}

func LoadConfig() model.Config {
	err := godotenv.Load()
	if err != nil {
		gologger.Fatal().Msg("Error loading .env file")
	}

	cfg := model.Config{}

	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPass = os.Getenv("DB_PASS")
	cfg.DBPort = convert(os.Getenv("DB_PORT"))
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.SSLMode = os.Getenv("SSL_MODE")
	cfg.Port = os.Getenv("PORT")

	return cfg
}

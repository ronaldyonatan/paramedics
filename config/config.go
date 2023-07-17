package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

var AppConfig Config

type Config struct {
	Smtp     Smtp
	Email    Email
	Postgres Postgres
	App      App
	Session  Session
}

type Smtp struct {
	Host string `env:"CONFIG_SMTP_HOST"`
	Post string `env:"CONFIG_SMTP_PORT"`
}

type Email struct {
	Email    string `env:"CONFIG_AUTH_EMAIL"`
	Password string `env:"CONFIG_AUTH_PASSWORD"`
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	DbName   string `env:"POSTGRES_DBNAME"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	SSLMode  string `env:"POSTGRES_SSLMODE"`
}
type App struct {
	BaseUrl     string `env:"BASE_URL"`
	BasePort    string `env:"BASE_PORT"`
	BaseWebPort string `env:"WEB_BASE_PORT"`
}
type Session struct {
	AuthSessionId string `env:"AUTH_SESSION"`
}

func LoadConfig() (cfg Config, err error) {
	err = godotenv.Load("../../.env")
	if err != nil {
		return
	}
	err = env.Parse(&cfg)
	return

}

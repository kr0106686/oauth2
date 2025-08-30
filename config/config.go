package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type (
	// Config -.
	Config struct {
		HTTP     HTTP
		DB       DB
		GRPC     GRPC
		JWT      JWT
		Provider Provider
	}

	// HTTP -.
	HTTP struct {
		Port string `env:"HTTP_PORT,required"`
	}

	// GRPC -.
	GRPC struct {
		Port string `env:"GRPC_PORT,required"`
	}

	// DB -.
	DB struct {
		Host string `env:"DB_HOST,required"`
		User string `env:"DB_USER,required"`
		Pass string `env:"DB_PASS,required"`
		Name string `env:"DB_NAME,required"`
		Port string `env:"DB_PORT,required"`
	}

	JWT struct {
		Secret string `env:"JWT_SECRET,required"`
	}

	Provider struct {
		Kakao  Kakao
		Google Google
	}

	Google struct {
		ClientID     string `env:"GOOGLE_CLIENT_ID"`
		ClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
		RedirectURI  string `env:"GOOGLE_REDIRECT_URI"`
	}

	Kakao struct {
		ClientID     string `env:"KAKAO_CLIENT_ID"`
		ClientSecret string `env:"KAKAO_CLIENT_SECRET"`
		RedirectURI  string `env:"KAKAO_REDIRECT_URI"`
	}
)

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

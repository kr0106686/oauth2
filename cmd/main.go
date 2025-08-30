package main

import (
	"log"

	"github.com/kr0106686/oauth2/v2/config"
	"github.com/kr0106686/oauth2/v2/internal/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("%v", err)
	}

	app.Run(cfg)
}

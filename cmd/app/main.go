package main

import (
	"log"

	"github.com/svy4toyboriy/work/config"
	"github.com/svy4toyboriy/work/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}

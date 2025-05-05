package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sonht1109/supercoder-go/internal/agent"
	"github.com/sonht1109/supercoder-go/internal/config"
	"github.com/sonht1109/supercoder-go/internal/global"
	"github.com/sonht1109/supercoder-go/internal/ui"
)

var AppConfig any

func main() {

	godotenv.Load() // use for development

	config.NewConfig()

	if global.Cfg.Model == "" {
		log.Fatal("Model is not set in the configuration")
	}

	agent := agent.NewCoderAgent("", global.Cfg.Model)
	ui.Run(agent.ChatAgent)
}

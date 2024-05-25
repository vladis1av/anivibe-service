package main

import (
	"anivibe-service/internal/config"
	"anivibe-service/internal/http"
	"anivibe-service/internal/infrastructure/routers/http/api"
	"context"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	log.Print("Init APP")

	mainRouter := mux.NewRouter()

	log.Print("Load Config")
	cfg := config.LoadConfig()

	httpConfig := &http.Config{
		ListenAddr:     cfg.HTTPAddr,
		IdleTimeout:    cfg.IdleTimeout,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		AllowedOrigins: cfg.AllowedOrigins,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
	}
	log.Printf("HTTP Config: %+v", httpConfig)

	log.Printf("Setup api routers ")
	api.SetupAPIRouters(mainRouter, cfg.AdminUserID)

	// log.Printf("Init telegram bot")
	// telegram.InitBot(cfg.TelegramBotToken, false)

	log.Print("Init Servers")
	mainServer := http.NewHTTP("ANIVIBE", mainRouter, httpConfig)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := mainServer.RunAndManageServers(ctx, mainServer); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

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
	proxyRouter := mux.NewRouter()

	api.SetupAPIRouters(mainRouter, proxyRouter)

	log.Print("Load Config")
	cfg := config.LoadConfig()

	httpConfig := &http.Config{
		ListenAddr:     cfg.HTTPAddr,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		IdleTimeout:    cfg.IdleTimeout,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
	}

	proxyConfig := &http.Config{
		ListenAddr:     cfg.ProxyHTTPAddr,
		ReadTimeout:    cfg.ProxyReadTimeout,
		WriteTimeout:   cfg.ProxyWriteTimeout,
		IdleTimeout:    cfg.ProxyIdleTimeout,
		MaxHeaderBytes: cfg.ProxyMaxHeaderBytes,
	}

	log.Print("Init Servers")
	mainServer := http.NewHTTP("ANIVIBE", mainRouter, httpConfig)
	proxyServer := http.NewHTTP("ANIVIBE_PROXY", proxyRouter, proxyConfig)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := mainServer.RunAndManageServers(ctx, mainServer, proxyServer); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

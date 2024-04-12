package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Config struct {
	ListenAddr         string
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleTimeout        time.Duration
	MaxHeaderBytes     int
	ServerWriteTimeout time.Duration
}

type Http struct {
	appName string
	server  *http.Server
	config  *Config
	quit    chan os.Signal
}

func NewHTTP(appName string, router *mux.Router, config *Config) *Http {
	server := &http.Server{
		Handler:        router,
		Addr:           config.ListenAddr,
		ReadTimeout:    config.ReadTimeout,
		IdleTimeout:    config.IdleTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: config.MaxHeaderBytes,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	return &Http{appName: appName, server: server, config: config, quit: quit}
}

func (s *Http) Run() {
	log.Printf("%s HTTP server running on %s", s.appName, s.config.ListenAddr)
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start %s  HTTP server: %v", s.appName, err)
	}
}

func (s *Http) Stop() {
	<-s.quit

	ctx, cancel := context.WithTimeout(context.Background(), s.config.ServerWriteTimeout)
	defer cancel()

	log.Printf("Shutting down %s HTTP server", s.appName)
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to gracefully shutdown %s HTTP server: %v", s.appName, err)
	}

	log.Printf("%s HTTP server stopped", s.appName)
	os.Exit(0)
}

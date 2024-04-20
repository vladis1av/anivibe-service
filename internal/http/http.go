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
	MaxHeaderBytes int
	ListenAddr     string
	AllowedOrigins []string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
}

type Http struct {
	appName string
	server  *http.Server
	config  *Config
}

func NewHTTP(appName string, router *mux.Router, config *Config) *Http {
	router.Use(corsMiddleware(config.AllowedOrigins))

	server := &http.Server{
		Handler:        router,
		Addr:           config.ListenAddr,
		ReadTimeout:    config.ReadTimeout,
		IdleTimeout:    config.IdleTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: config.MaxHeaderBytes,
	}

	return &Http{appName: appName, server: server, config: config}
}

func (s *Http) Run() error {
	log.Printf("%s HTTP server running on %s", s.appName, s.config.ListenAddr)
	return s.server.ListenAndServe()
}

func (s *Http) Stop(ctx context.Context) error {
	log.Printf("Shutting down %s HTTP server", s.appName)
	return s.server.Shutdown(ctx)
}

func (s *Http) RunAndManageServers(ctx context.Context, servers ...*Http) error {
	errChan := make(chan error, len(servers))

	for _, server := range servers {
		go func(s *Http) {
			errChan <- s.Run()
		}(server)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	select {
	case <-ctx.Done():
		return nil
	case <-quit:
		for _, server := range servers {
			if err := server.Stop(ctx); err != nil {
				log.Printf("Failed to stop %s HTTP server: %v", server.appName, err)
			}
		}
		return nil
	case err := <-errChan:
		if err != nil && err != http.ErrServerClosed {
			return err
		}
	}

	return nil
}

func corsMiddleware(allowedOrigins []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Получаем источник запроса
			origin := r.Header.Get("Origin")

			// Проверяем, разрешен ли источник
			allowed := false
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					allowed = true
					break
				}
			}

			if allowed {
				// Разрешаем источник
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			// Разрешить заголовки
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			// Разрешить методы
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

			// Обработать предварительные запросы OPTIONS
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

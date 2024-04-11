package app

import (
	v1 "anivibe-service/internal/infrastructure/routers/http/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Http struct {
	listenAddr string
}

func NewApp(listenAddr string) *Http {
	return &Http{listenAddr: listenAddr}
}

func (s *Http) Run() {
	router := mux.NewRouter()

	v1.SetupAPIRouters(router)

	log.Println("HTTP server running on port: ", s.listenAddr)

	log.Fatal(http.ListenAndServe(s.listenAddr, router))
}

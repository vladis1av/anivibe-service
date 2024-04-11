package v1

import (
	v1 "anivibe-service/internal/infrastructure/routers/http/api/v1"
	"log"

	"github.com/gorilla/mux"
)

func SetupAPIRouters(router *mux.Router) {
	log.Println("setup routers")

	apiRouter := router.PathPrefix("/api").Subrouter()
	v1Router := apiRouter.PathPrefix("/v1").Subrouter()

	v1.SetupV1ImageRouters(v1Router)
	v1.SetupV1MangaRouters(v1Router)
}

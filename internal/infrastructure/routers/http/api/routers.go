package api

import (
	v1 "anivibe-service/internal/infrastructure/routers/http/api/v1"
	"log"

	"github.com/gorilla/mux"
)

const prefixV1 = "/api/v1"

func SetupAPIRouters(mainRouter *mux.Router) {
	log.Println("setup routers")

	mainRouterV1 := mainRouter.PathPrefix(prefixV1).Subrouter()

	v1.SetupV1MangaRouters(mainRouterV1)
	v1.SetupV1ImageRouters(mainRouterV1)
}

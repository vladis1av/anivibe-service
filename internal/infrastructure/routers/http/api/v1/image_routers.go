package v1

import (
	"anivibe-service/internal/infrastructure/handlers/http/v1/image"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupV1ImageRouters(router *mux.Router) {
	log.Println("setup image routers v1")

	imageRouter := router.PathPrefix("/image").Subrouter()

	imageRouter.HandleFunc("/proxy", image.Proxy).Methods(http.MethodGet)
	imageRouter.HandleFunc("/proxyyy", image.Proxy).Methods(http.MethodGet)
}

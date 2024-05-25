package v1

import (
	"anivibe-service/internal/infrastructure/handlers/http/v1/sitemap"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupV1SitemapRouters(router *mux.Router) {
	log.Println("setup sitemap routers v1")

	imageRouter := router.PathPrefix("/sitemap").Subrouter()

	imageRouter.HandleFunc("", sitemap.GetSitemap).Methods(http.MethodGet)
}

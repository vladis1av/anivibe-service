package v1

import (
	"anivibe-service/internal/infrastructure/handlers/http/v1/manga"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupV1MangaRouters(router *mux.Router) {
	log.Println("setup manga routers v1")

	mangaRouter := router.PathPrefix("/manga").Subrouter()
	mangaRouter.StrictSlash(true)

	mangaRouter.HandleFunc("", manga.GetMangas).Methods(http.MethodGet)
	mangaRouter.HandleFunc("/{id}", manga.GetMangaById).Methods(http.MethodGet)
	mangaRouter.HandleFunc("/{mangaId}/chapter/{chapterId}", manga.GetMangaChapterById).Methods(http.MethodGet)
}

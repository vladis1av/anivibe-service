package manga

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/vladis1av/desume-client-go/desume"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json"
)

const timeout = 5 * time.Second

type response struct {
	PageNavParams *desume.PageNavParams `json:"pageNavParams,omitempty"`
	Response      interface{}           `json:"response,omitempty"`
	Error         string                `json:"error,omitempty"`
}

func sendResponse(w http.ResponseWriter, data response, statusCode int, err error) {
	resp := response{
		PageNavParams: data.PageNavParams,
		Response:      data.Response,
	}

	if err != nil {
		errResp := desume.MangaError{Error: err.Error()}
		w.Header().Set(contentType, applicationJSON)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(errResp)
		return
	}

	w.Header().Set(contentType, applicationJSON)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

// TODO Вынести в конфиг и прокидывать сюда параметры
var client = desume.NewClient()

func GetMangaById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid manga ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	manga, err := client.GetMangaById(ctx, int(id))
	if err != nil {
		sendResponse(w, response{Response: nil, PageNavParams: nil}, http.StatusInternalServerError, err)
		return
	}

	sendResponse(w, response{Response: manga.Response}, http.StatusOK, nil)
}

func GetMangaChapterById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mangaIdStr := vars["mangaId"]
	mangaChapterIdStr := vars["chapterId"]

	baseRes := response{Response: nil, PageNavParams: nil}

	mangaId, err := strconv.Atoi(mangaIdStr)
	if err != nil {
		sendResponse(w, baseRes, http.StatusBadRequest, fmt.Errorf("Invalid manga ID: %v", err))
		return
	}

	mangaChapter, err := strconv.Atoi(mangaChapterIdStr)
	if err != nil {
		sendResponse(w, baseRes, http.StatusBadRequest, fmt.Errorf("Invalid manga chapter ID: %v", err))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()
	manga, err := client.GetMangaChapter(ctx, mangaId, mangaChapter)
	if err != nil {
		sendResponse(w, baseRes, http.StatusInternalServerError, err)
		return
	}

	baseRes.Response = manga.Response

	sendResponse(w, baseRes, http.StatusOK, nil)
}

func GetMangas(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	order := r.URL.Query().Get("order")
	kinds := r.URL.Query().Get("kinds")
	genres := r.URL.Query().Get("genres")
	search := r.URL.Query().Get("search")

	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	filterParams := desume.GetMangaFilterParams(page, limit, order, kinds, genres, search)
	mangasFiltered, err := client.GetMangas(ctx, filterParams)
	if err != nil {
		sendResponse(w, response{PageNavParams: nil}, http.StatusInternalServerError, err)
		return
	}

	sendResponse(w, response{
		Response:      mangasFiltered.Response,
		PageNavParams: &mangasFiltered.PageNavParams,
	}, http.StatusOK, nil)
}

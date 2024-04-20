package image

import (
	"anivibe-service/internal/utils"
	"io"
	"log"
	"net/http"
)

func Proxy(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	referer := r.URL.Query().Get("referer")
	authority := r.URL.Query().Get("authority")

	imageBody, err := utils.FetchImage(url, referer, authority)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer imageBody.Close()

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK) // Установка кода статуса перед io.Copy
	_, err = io.Copy(w, imageBody)
	if err != nil {
		log.Printf("Error sending image from %s: %v", url, err)
		return
	}
}

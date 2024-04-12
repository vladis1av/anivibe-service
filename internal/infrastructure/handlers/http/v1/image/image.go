package image

import (
	"anivibe-service/internal/utils"
	"io"
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

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	io.Copy(w, imageBody)
}

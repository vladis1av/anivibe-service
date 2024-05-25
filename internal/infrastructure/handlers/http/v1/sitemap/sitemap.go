package sitemap

import (
	"io"
	"net/http"
	"os"
)

const (
	contentType    = "Content-Type"
	applicationXML = "application/xml"
)

func GetSitemap(w http.ResponseWriter, r *http.Request) {
	version := r.URL.Query().Get("version") // old or new - default new
	const oldPath = "/data/vercel/sitemap.xml"
	const currentPath = "/data/sitemap.xml"

	var filePath string
	if version == "old" {
		filePath = oldPath
	} else {
		filePath = currentPath
	}

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set(contentType, applicationXML)

	w.WriteHeader(http.StatusOK)
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

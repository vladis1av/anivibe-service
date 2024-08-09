package utils

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func FetchImage(url string, referer string, authority string) (io.ReadCloser, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil // Разрешить все редиректы
		},
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if referer != "" {
		req.Header.Set("Referer", referer)
	}

	if authority == "" {
		authority = GetAuthorityFromUrl(url)
	}

	req.Header.Set("Authority", authority)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to fetch image: %s", resp.Status)
	}

	return resp.Body, nil
}

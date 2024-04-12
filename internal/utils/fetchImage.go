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

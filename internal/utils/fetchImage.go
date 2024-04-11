package utils

import (
	"io"
	"net/http"
)

func FetchImage(url string, referer string, authority string) ([]byte, error) {
	var bodyText []byte

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return bodyText, err
	}

	if authority == "" {
		authority = GetAuthorityFromUrl(url)
	}

	req.Header.Set("authority", authority)
	req.Header.Set("accept", "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8")
	req.Header.Set("accept-language", "ru,en;q=0.9")
	req.Header.Set("referer", referer)
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "YaBrowser";v="24.1", "Yowser";v="2.5"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "image")
	req.Header.Set("sec-fetch-mode", "no-cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 YaBrowser/24.1.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return bodyText, err
	}

	defer resp.Body.Close()

	bodyText, err = io.ReadAll(resp.Body)
	if err != nil {
		return bodyText, err
	}

	return bodyText, nil
}

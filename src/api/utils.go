package api

import (
    "net/http"
)

// SetCommonHeaders sets common headers for API requests
func SetCommonHeaders(req *http.Request, sessionKey string) {
    req.Header.Set("authority", "apiv3.fansly.com")
    req.Header.Set("accept", "application/json, text/plain, */*")
    req.Header.Set("accept-language", "en;q=0.8,en-US;q=0.7")
    req.Header.Set("authorization", sessionKey)
    req.Header.Set("origin", "https://fansly.com")
    req.Header.Set("referer", "https://fansly.com/")
    req.Header.Set("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"Google Chrome\";v=\"114\"")
    req.Header.Set("sec-ch-ua-mobile", "?0")
    req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
    req.Header.Set("sec-fetch-dest", "empty")
    req.Header.Set("sec-fetch-mode", "cors")
    req.Header.Set("sec-fetch-site", "same-site")
    req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
}
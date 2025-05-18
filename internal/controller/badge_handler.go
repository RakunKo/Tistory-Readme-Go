package controller

import (
	"net/http"
	"strings"
	"tistory-readme-go/internal/service"
)

func BadgeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		http.Error(w, "userName not specified", http.StatusBadRequest)
		return
	}
	userName := parts[2]

	channel, err := service.GetTistoryRss(userName)
	if err != nil {
		http.Error(w, "Failed to fetch RSS: "+err.Error(), http.StatusInternalServerError)
		return
	}

	svg := service.BuildSvg(channel)
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(svg))
}

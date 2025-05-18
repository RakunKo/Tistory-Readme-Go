package handler

import (
	"github.com/RakunKo/Tistory-Readme-Go/internal/controller"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	controller.BadgeHandler(w, r)
}

package main

import (
	"fmt"
	"github.com/RakunKo/Tistory-Readme-Go/pkg/controller"
	"log"
	"net/http"
)

/*
*
Local에서 테스트시에 활용합니다.
path : tistory-badge/{userName}/{color}
color에서 #은 제외
*/
func main() {
	http.HandleFunc("/tistory-badge/", controller.BadgeHandler)
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

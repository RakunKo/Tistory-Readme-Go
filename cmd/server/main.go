package main

import (
	"fmt"
	"log"
	"net/http"
	"tistory-readme-go/internal/controller"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func Handler() {
	http.HandleFunc("/tistory-badge/", controller.BadgeHandler)
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

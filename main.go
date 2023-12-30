package main

import (
	"fmt"
	"log"
	"net/http"
	"one_pane_assessment/handler"
)

func main() {
	h := handler.NewHandler()
	http.HandleFunc("/", h.HandleRoute)
	fmt.Println("server running on port:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("server failed to start")
	}
}

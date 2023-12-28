package main

import (
	"net/http"
	"one_pane_assessment/handler"
)

func main() {
	h := handler.NewHandler()
	http.HandleFunc("/", h.HandleRoute)
	http.ListenAndServe(":8080", nil)
}

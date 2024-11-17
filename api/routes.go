package api

import (
	"image_filters/api/handlers"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}

func SetupRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /", helloWorld)
	router.HandleFunc("GET /convert/ascii", handlers.AsciiHandler)
}

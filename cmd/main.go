package main

import (
	"image_filters/api"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	api.SetupRoutes(router)

	log.Println("Server listening to http://127.0.0.1:8000")

	http.ListenAndServe(":8000", router)
}

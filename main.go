package main

import (
	"goserver/rest"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	http.HandleFunc("/", rest.Welcome)
	chi := chi.NewRouter()
	http.ListenAndServe(":8080", nil)
}

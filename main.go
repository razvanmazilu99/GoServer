package main

import (
	"goserver/config"
	"goserver/db"
	"goserver/rest"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Not able to load env file")
	}

	err = config.InitConfig()

	if err != nil {
		log.Fatal("Not able to create config")
	}

	db.Initdatabase("postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable", 1)

	http.HandleFunc("/", rest.Welcome)
	http.HandleFunc("/health", rest.Welcome)
	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		r.Get("/person", rest.GetPerson)
		r.Post("/person", rest.PostPerson)
		r.Delete("/person", rest.DeletePerson)
	})
	http.ListenAndServe(":8080", router)
}

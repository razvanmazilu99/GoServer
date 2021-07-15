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

	db.Initdatabase()
	db.InitData()

	var endpoint = "/person"

	http.HandleFunc("/", rest.Welcome)
	http.HandleFunc("/health", rest.Welcome)
	router := chi.NewRouter()
	router.Route("/"+config.GetConfig().APPVersion, func(r chi.Router) {
		r.Get(endpoint, rest.GetPerson)
		r.Post(endpoint, rest.PostPerson)
		r.Delete(endpoint, rest.DeletePerson)
		r.Post("/login", rest.Login)
		r.Get("/welcome", rest.Welcome1)
		r.Get("/logout", rest.Logout)
	})
	http.ListenAndServe(":"+config.GetConfig().Port, router)
}

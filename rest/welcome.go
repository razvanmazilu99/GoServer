package rest

import "net/http"

func Welcome(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Welcome to Atoss"))
}

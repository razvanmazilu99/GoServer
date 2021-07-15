package rest

import (
	"encoding/json"
	"fmt"
	"goserver/entity"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func Login(rw http.ResponseWriter, req *http.Request) {

	body := req.Body
	bodyBytes, _ := ioutil.ReadAll(body) //we should handle the error in here
	var credentials entity.Credential

	json.Unmarshal(bodyBytes, &credentials)
	session, _ := store.Get(req, "cookie-name")
	session.ID = credentials.ID
	session.Values["userID"] = credentials.ID
	session.Values["authenticated"] = true
	fmt.Fprintln(rw, "Login Successful")
	session.Save(req, rw)
}

func Welcome1(rw http.ResponseWriter, req *http.Request) {

	name := req.URL.Query().Get("name")
	session, _ := store.Get(req, "cookie-name")
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(rw, "Access forbidden", http.StatusForbidden)
		return
	}
	fmt.Fprintln(rw, "Welcome", name)
}

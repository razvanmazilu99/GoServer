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

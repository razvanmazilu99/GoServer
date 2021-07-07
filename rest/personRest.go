package rest

import (
	"encoding/json"
	"fmt"
	"goserver/entity"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func PostPerson(rw http.ResponseWriter, r *http.Request) {

	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	var person entity.Person
	err = json.Unmarshal(bodyBytes, &person)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	fmt.Println(person)
	rw.Write(bodyBytes)
}

func GetPerson(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	rw.Write([]byte(name))
}

func hasError(rw http.ResponseWriter, err error, message string) bool {
	logger := new(logrus.Entry)

	if err != nil {
		logger.WithError(err).Error(message)
		rw.Write([]byte(message))
		return true
	}

	return false
}

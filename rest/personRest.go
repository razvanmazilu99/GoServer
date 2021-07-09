package rest

import (
	"encoding/json"
	"fmt"
	"goserver/db"
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

	db.GetDB().Create(&person)

	fmt.Println(person)
	rw.Write(bodyBytes)
}

func GetPerson(rw http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("id")

	var person entity.Person

	result := db.GetDB().Where("id=?", name).Find(&person)

	if result.RecordNotFound() {
		http.Error(rw, "No record", http.StatusInternalServerError)
		return
	}

	if result.Error != nil {
		http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
	}

	personBytes, err := json.Marshal(person)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	rw.Write(personBytes)
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

func DeletePerson(rw http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	result := db.GetDB().Delete(&entity.Person{}, "id=?", id)

	if result.Error != nil {
		http.Error(rw, "Internal error. Please try again after a while", http.StatusInternalServerError)
		return
	}
	rw.Write([]byte("Record successfully"))
}

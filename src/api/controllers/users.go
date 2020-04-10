package controllers

import (
	"api/models"
	"api/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UserController struct {
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("List Users"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List User"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}

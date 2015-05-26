package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devin0910/go-rest/models"
	"github.com/julienschmidt/httprouter"
)

type (
	UserController struct{}
)

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "Bob",
		Gender: "male",
		Age:    50,
		Id:     p.ByName("id"),
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = "foo"

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(200)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devin0910/go-rest/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	r.GET("/user/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		u := models.User{
			Name:   "Bob Smith",
			Gender: "male",
			Age:    50,
			Id:     p.ByName("id"),
		}

		uj, _ := json.Marshal(u)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", uj)
	})

	r.POST("/user", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		u := models.User{}

		json.NewDecoder(r.Body).Decode(&u)

		u.Id = "foo"

		uj, _ := json.Marshal(u)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", uj)
	})

	r.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome!\n")
	})

	http.ListenAndServe("localhost:3000", r)
}

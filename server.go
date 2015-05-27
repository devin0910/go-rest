package main

import (
	"fmt"
	"net/http"

	"github.com/devin0910/go-rest/controllers"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.RemoveUser)

	r.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome!\n")
	})

	http.ListenAndServe("localhost:3000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://121.40.58.115")

	if err != nil {
		panic(err)
	}
	return s
}

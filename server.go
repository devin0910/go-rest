package main

import (
	"fmt"
	"net/http"

	"github.com/devin0910/go-rest/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController()

	r.GET("/user/:id", uc.GetUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.RemoveUser)

	r.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome!\n")
	})

	http.ListenAndServe("localhost:3000", r)
}

package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
	models "../models"
	_ "github.com/satori/go.uuid"
	"time"
)

func handleMain(w http.ResponseWriter,r *http.Request) {
	t,err := template.ParseFiles("../templates/index.html")
	if err != nil {
		fmt.Fprintf(w,"500 Internal server error")
		return
	}
	t.Execute(w,nil)
}

func handleRegistration(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var user models.User
		var err error
		id, err := uuid.NewV4()
		if err != nil {
			fmt.Fprintf(w, "%v Internal server error", http.StatusInternalServerError)
			return
		}

		user.Username = r.FormValue("username")
		user.Password = r.FormValue("password")
		user.Email = r.FormValue("email")
		user.Id = id.String()
		user.RegistrationDate = time.Now().String()

		err = Register(user)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		fmt.Println(user)

	case "GET":
		t,err := template.ParseFiles("../templates/registration.html")
		if err != nil {
			fmt.Fprintf(w,"500 Internal server error")
			return
		}
		t.Execute(w,nil)
	}

}
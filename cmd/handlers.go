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
	response, status := authenticate(r)
	if status != http.StatusOK{
		w.WriteHeader(status)
		return
	}

	fmt.Fprintf(w,"Welcome %s",response)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t,err := template.ParseFiles("../templates/authentication.html")
		if err != nil {
			fmt.Fprintf(w,"500 Internal server error")
			return
		}
		t.Execute(w,nil)

	case "POST":
		username := r.FormValue("username")
		password := r.FormValue("password")
		err := correctUser(username,password)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		sessionToken, _ := uuid.NewV4()
		cache[sessionToken.String()] = username
		http.SetCookie(w, &http.Cookie{
			Name : "session_token",
			Value : sessionToken.String(),
			Expires: time.Now().Add(120 * time.Second),
			HttpOnly: true,
		})
		fmt.Fprintf(w,"Welcome!")

	default:
		fmt.Fprintf(w,"400 bad request")
	}
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

	default:
		fmt.Fprintf(w,"400 bad request")
	}
}

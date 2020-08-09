package main

import (
	models "../models"
	"fmt"
	_ "github.com/satori/go.uuid"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
	"time"
)


func handleMain(w http.ResponseWriter,r *http.Request) {
	t,err := template.ParseFiles("../templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.Execute(w,posts.Body)
}

func writePost(w http.ResponseWriter, r *http.Request){
	t,err := template.ParseFiles("../templates/write.html")

	_ , status := authenticate(r)
	if status != http.StatusOK{
		http.Redirect(w,r,"/authentication",status)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "write",nil)
}

func savepostHandler(w http.ResponseWriter, r *http.Request){
	var post models.Post
	var err error

	post.Id = GenerateId()
	post.Description = r.FormValue("description")
	t := time.Now()
	post.PostDate = t.Format(time.RFC1123)
	userid , status := authenticate(r)
	if status != http.StatusOK{
		http.Redirect(w,r,"/authentication",status)
		return
	}
	post.UserId = userid
	post.Category = "choumi"
	post.Theme = r.FormValue("theme")

	err = NewPost(post)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	http.Redirect(w,r,"/", 302)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t,err := template.ParseFiles("../templates/authentication.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w,"%v",http.StatusInternalServerError)
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
			Expires: time.Now().Add(1 * time.Hour),
			HttpOnly: true,
		})
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return

	default:
		w.WriteHeader(http.StatusBadRequest)
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
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "500 Internal server error")
			return
		}

		user.Username = r.FormValue("username")
		user.Password = r.FormValue("password")
		user.Email = r.FormValue("email")
		user.Id = id.String()
		user.RegistrationDate = time.Now().String()

		err = register(user)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}

	case "GET":
		t,err := template.ParseFiles("../templates/registration.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w,"500 Internal server error")
			return
		}
		t.Execute(w,nil)

	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"400 bad request")
	}
}

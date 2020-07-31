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


var posts map[string]*models.Post

func handleMain(w http.ResponseWriter,r *http.Request) {
	response, status := authenticate(r)
	if status != http.StatusOK{
		t,err := template.ParseFiles("../templates/index.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(status)
		t.Execute(w,nil)
		return
	} else {
		t,err := template.ParseFiles("../templates/index.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(status)
		t.Execute(w,posts)
	}
	fmt.Println(posts)
	fmt.Println("posts")
	_ = response
}

func writePost(w http.ResponseWriter, r *http.Request){
	t,err := template.ParseFiles("../templates/write.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "write",nil)
}

func savepostHandler(w http.ResponseWriter, r *http.Request){

	id := GenerateId()
	description := r.FormValue("description")

	t := time.Now()
	postdate := t.Format(time.RFC1123)

	userid , _ :=  authenticate(r)
	category := ""
	theme := r.FormValue("theme")

	post := models.NewPost(id, description, postdate, userid, category, theme)
	posts[post.Id] = post
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
		fmt.Fprintf(w,"Welcome!")

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

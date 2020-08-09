package main

import (
	models "../models"
	"fmt"
	"log"
	"net/http"
)

var cache map[string]string

var users models.Users
var posts models.Posts

func main() {
	err := models.Init("forum.db")
	if err != nil {
		log.Fatal(err)
	}
	cache = make(map[string]string)

	//css/
	css := http.FileServer(http.Dir("css"))

	// init models
	err = users.Init()
	if err != nil{
		log.Fatal(err.Error())
	}

	err = posts.Init()
	if err != nil{
		log.Fatal(err)
	}

	//routes
	http.Handle("/css/", http.StripPrefix("/css/", css))
	http.HandleFunc("/write",writePost)
	http.HandleFunc("/",handleMain)
	http.HandleFunc("/savePost",savepostHandler)
	http.HandleFunc("/registration",handleRegistration)
	http.HandleFunc("/authentication",handleAuth)

	fmt.Println("hi")
	log.Fatal(http.ListenAndServe(":3030", nil))
}

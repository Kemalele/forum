package main

import (
	models "../models"
	"fmt"
	"log"
	"net/http"
	router "../pkg/router"
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
	//css := http.FileServer(http.Dir("css"))

	// init models
	err = users.Init()
	if err != nil{
		log.Fatal(err.Error())
	}

	err = posts.Init()
	if err != nil{
		log.Fatal(err)
	}

	r := router.New(getMain)
	r.Handle("GET","/",getMain)
	r.Handle("GET","/write",writePost)
	r.Handle("GET","/registration",getRegistration)
	r.Handle("GET","/authentication",getAuth)
	r.Handle("GET","/post/:id",handlePostPage)

	//r.Handle("POST","/",handleMain)
	r.Handle("POST","/savePost",savepostHandler)
	r.Handle("POST","/registration",handleRegistration)
	r.Handle("POST","/authentication",handleAuth)

	//routes
	//http.Handle("/css/", http.StripPrefix("/css/", css))
	//http.HandleFunc("/write",writePost)
	//http.HandleFunc("/",handleMain)
	//http.HandleFunc("/post/:id",handlePostPage)
	//http.HandleFunc("/savePost",savepostHandler)
	//http.HandleFunc("/registration",handleRegistration)
	//http.HandleFunc("/authentication",handleAuth)

	fmt.Println("hi")
	log.Fatal(http.ListenAndServe(":3030", r))
}

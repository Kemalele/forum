package main

import (
	models "../models"
	"fmt"
	"log"
	"net/http"
)

var cache map[string]string

func main() {
	fmt.Println("hi")

	err := models.Init("forum.db")
	if err != nil {
		log.Fatal(err)
	}
	cache = make(map[string]string)
	//css/
	css := http.FileServer(http.Dir("css"))
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/css/", http.StripPrefix("/css/", css))

	posts = make(map[string]*models.Post, 0)

	http.HandleFunc("/write",writePost)
	http.HandleFunc("/",handleMain)
	http.HandleFunc("/savePost",savepostHandler)
	http.HandleFunc("/registration",handleRegistration)
	http.HandleFunc("/authentication",handleAuth)
	log.Fatal(http.ListenAndServe(":3030", nil))

}

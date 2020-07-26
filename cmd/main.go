package main

import (
	models "../models"
	"log"
	"net/http"

)

var cache map[string]string

func main() {
	err := models.Init("forum.db")
	if err != nil {
		log.Fatal(err)
	}
	cache = make(map[string]string)

	http.HandleFunc("/",handleMain)
	http.HandleFunc("/registration",handleRegistration)
	http.HandleFunc("/authentication",handleAuth)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

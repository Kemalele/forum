package main

import (
	//"log"
	//"net/http"
)

import (
	models "../models"
	"log"
	"net/http"

	//"log"
)

func main() {
	err := models.Init("forum.db")
	if err != nil {
		log.Fatal(err)
	}
	//users,posts,comments,likes := InitModels()
	http.HandleFunc("/",handleMain)
	http.HandleFunc("/registration",handleRegistration)
	http.HandleFunc("/authentication",handleAuth)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

//func InitModels() (models.Users,models.Posts,models.Comments,models.Likes){
//	var usrs models.Users
//	var posts models.Posts
//	var comments models.Comments
//	var likes models.Likes
//
//	usrs.Init()
//	posts.Init()
//	comments.Init()
//	likes.Init()
//
//	return usrs,posts,comments,likes
//}

package main

import (
	models "../models"
	"fmt"

	//"log"
)
func main() {
	models.Init("forum.db")
	//var users models.Users
	//users.Add(models.User{"1","1","1","1","1"},models.Db)
	//fmt.Println(users)

	//var posts models.Posts
	//err := posts.Add(models.Post{"2","desc","01-01-2020","1"},models.Db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(posts)

	var comms models.Comments
	comms.Init()
	comms.Add(models.Comment{"1","desc","01-01-2020","3","10"},models.Db)
	fmt.Println(comms)
}
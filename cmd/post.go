package main
import (
	models "../models"
	"errors"
	"fmt"
)

func NewPost(post models.Post)  error {
	var posts models.Posts
	err := posts.Init()
	fmt.Println(err.Error())

	err = validPost(post)
	fmt.Println(post)
	if err != nil {
		return err
	}

	posts.Add(post,models.Db)
	return nil
}

func validPost(p models.Post) error{
	if len(p.Theme) > 1 {
		return errors.New("title must be at least 1 symbol")
	}

	if len(p.Description) > 1 {
		return errors.New("content must be at least 1 symbol")
	}

	return nil
}
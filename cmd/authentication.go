package main

import (
	models "../models"
	"errors"
	"net/http"
)

func correctUser(username,password string) error{
	var users models.Users
	users.Init()
	for _, usr := range users.Body {
		if usr.Username == username {
			decryptedPass, _ := decrypt(usr.Password)
			if decryptedPass == password {
				return nil
			}else {
				return errors.New("wrong password")
			}
		}
	}
	return errors.New("wrong username")
}

func authenticated(r *http.Request) (string,bool){
	c,err := r.Cookie("session_token")
	if err != nil {
		return "", false
	}

	sessionToken := c.Value

	// nickname
	response := cache[sessionToken]
	if response == "" {
		return "", false
	}

	return response,true
}

//func authenticated(r *http.Request) bool {
//	_,err := r.Cookie("session_token")
//	if err != nil {
//		return false
//	}
//
//	return true
//}
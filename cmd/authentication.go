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

func authenticate(r *http.Request) (string,int){
	c,err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return "",http.StatusUnauthorized
		}
		return "",http.StatusBadRequest
	}

	sessionToken := c.Value
	response := cache[sessionToken]
	if response == "" {
		return "",http.StatusUnauthorized
	}

	return response,http.StatusOK
}
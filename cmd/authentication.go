package main

import (
	models "../models"
	"errors"
)

func Authenticate(username,password string) error{
	var users models.Users
	users.Init()
	for _, usr := range users.Body {
		if usr.Username == username {
			if usr.Password == password {
				return nil
			}else {
				return errors.New("wrong password")
			}
		}
	}

	return errors.New("wrong username")
}


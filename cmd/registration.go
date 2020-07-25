package main

import (
	models "../models"
	"errors"
)

func Register(usr models.User)  error {
	var users models.Users
	users.Init()

	err := validUser(usr)
	if err != nil {
		return err
	}
	err = uniqueUser(usr)
	if err != nil {
		return err
	}

	users.Add(usr,models.Db)
	return nil
}

func validUser(usr models.User) error{
	if len(usr.Password) < 6 {
		return errors.New("password must be at least 6 symbols")
	}

	if len(usr.Username) < 3 {
		return errors.New("username must be at least 3 symbols")
	}

	return nil
}

func uniqueUser(usr models.User) error{
	var users models.Users
	users.Init()

	for _,user := range users.Body {
		if user.Username == usr.Username{
			return errors.New("username already taken")
		}

		if user.Email == usr.Email {
			return errors.New("email already taken")
		}
	}

	return nil
}
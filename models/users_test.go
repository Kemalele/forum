package models
import "testing"

func TestUsers_GetUserById(t *testing.T) {
	var users Users
	var mdb MockDB
	users.Add(User{Id: "1",Username: "username",Password: "pass",Email: "email",RegistrationDate: "01-02-2019"}, &mdb)
	usr,err := users.GetUserById("1")

	if err != nil {
		t.Fatal("can't get user by id")
	}

	if usr.Id != "1" || usr.Username != "username" || usr.Password != "pass" {
		t.Fatal("can't get user by id")
	}
 }

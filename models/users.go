package models

type Users struct {
	Body []User
}

func (users *Users) Init() error {
	rows,err := Db.Query("SELECT * FROM User")
	if err != nil {
		return err
	}
	for rows.Next() {
		usr := User{}
		err := rows.Scan(&usr.Id,&usr.Username,&usr.Password,&usr.Email,&usr.RegistrationDate)
		if err != nil {
			return err
		}
		users.Body = append(users.Body,usr)
	}
	return nil
}

func (users *Users) Add(usr User,sql SQLDB) error{
	_,err := sql.Exec("INSERT INTO USER (Id,Username,Password,Email,RegistrationDate) values ($1,$2,$3,$4,$5)",usr.Id,usr.Username,usr.Password,usr.Email,usr.RegistrationDate)
	if err != nil {
		return err
	}
	users.Body = append(users.Body,usr)
	return nil
}

func (users *Users) GetUserById(Id string)(User,error) {
	for _,usr := range users.Body {
		if usr.Id == Id {
			return usr,nil
		}
	}
	return User{},nil
}

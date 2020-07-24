package models

type Likes struct {
	Body []Liked
}

func (l *Likes) Init() error {
	rows,err := Db.Query("SELECT * FROM Liked")
	if err != nil {
		return err
	}
	for rows.Next() {
		liked := Liked{}
		err := rows.Scan(&liked.Id,&liked.Value,&liked.PostId,&liked.UserId)
		if err != nil {
			return err
		}
		l.Body = append(l.Body,liked)
	}
	return nil
}

func (l *Likes) Add(liked Liked,sql SQLDB) error{
	_,err := sql.Exec("INSERT INTO LIKED (Id,Value,PostId,UserId) values ($1,$2,$3,$4)",liked.Id,liked.Value,liked.PostId,liked.UserId)
	if err != nil {
		return err
	}
	l.Body = append(l.Body,liked)
	return nil
}

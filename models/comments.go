package models

type Comments struct {
	Body []Comment
}

func (c *Comments) Init() error {
	rows,err := Db.Query("SELECT * FROM Comment")
	if err != nil {
		return err
	}
	for rows.Next() {
		comment := Comment{}
		err := rows.Scan(&comment.Id,&comment.Description,&comment.PostDate,&comment.UserId,&comment.PostId)
		if err != nil {
			return err
		}
		c.Body = append(c.Body,comment)
	}
	return nil
}

func (c *Comments) Add(comment Comment,sql SQLDB) error{
	_,err := sql.Exec("INSERT INTO COMMENT (Id,Description,Post_date,UserId,PostId) values ($1,$2,$3,$4,$5)",comment.Id,comment.Description,comment.PostDate,comment.UserId,comment.PostId)
	if err != nil {
		return err
	}
	c.Body = append(c.Body,comment)
	return nil
}
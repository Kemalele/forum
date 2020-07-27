package models

type Posts struct {
	Body []Post
}

func (p *Posts) Init() error {
	rows,err := Db.Query("SELECT * FROM Post")
	if err != nil {
		return err
	}
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.Id,&post.Description,&post.PostDate,&post.UserId,&post.Category)
		if err != nil {
			return err
		}
		p.Body = append(p.Body,post)
	}
	return nil
}

func (p *Posts) Add(post Post,sql SQLDB) error{
	_,err := sql.Exec("INSERT INTO POST (Id,Description,Post_date,UserId,Category) values ($1,$2,$3,$4,$5)",post.Id,post.Description,post.PostDate,post.UserId,post.Category)
	if err != nil {
		return err
	}
	p.Body = append(p.Body,post)
	return nil
}
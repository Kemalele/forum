package models

import (
	"errors"
	"fmt"
)

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
		err := rows.Scan(&post.Id,&post.Description,&post.PostDate,&post.UserId,&post.Category,&post.Theme)
		if err != nil {
			return err
		}
		p.Body = append(p.Body,post)
	}
	return nil
}

func (p *Posts) Add(post Post,sql SQLDB) error{
	_,err := sql.Exec("INSERT INTO POST (Id,Description,Post_date,UserId,Category,Theme) values ($1,$2,$3,$4,$5,$6)",post.Id,post.Description,post.PostDate,post.UserId,post.Category,post.Theme)
	if err != nil { 
		return err
	}
	p.Body = append(p.Body,post)
	fmt.Println(post)
	return nil
}

func (p *Posts) GetPostById(Id string)(Post,error) {
	for _,post := range p.Body {
		if post.Id == Id {
			return post,nil
		}
	}
	return Post{},errors.New("no such user")
}
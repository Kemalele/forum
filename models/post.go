package models

type Post struct {
	Id          string
	Description string
	PostDate    string
	UserId      string
	Category 	string
	Theme 		string
}

func NewPost(id, description, postdate, userid, category, theme string) *Post{
	return &Post{id, description, postdate, userid, category, theme}
}
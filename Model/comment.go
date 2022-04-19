package Model

type Comment struct {
	ID       int    `form  id json : id`
	Username string `form : username json : username`
	Email    string `form : email json : email`
	Comment  string `form : comment json : comment`
}

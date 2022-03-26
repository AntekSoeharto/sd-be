package Model

type Comment struct {
	ID      int    `form  id json : id`
	User    User   `form : userid json : userid`
	Comment string `form : comment json : cmment`
}

package Model

type User struct {
	ID       int    `form : id json : id`
	Name     string `form : name json : name`
	Username string `form : username json : username`
	Email    string `form : email json : email`
	Gender   string `form : gender json : gender`
	BirthDay string `form : birthday json : birthday`
	Password string `form : password json : password`
}

package Model

type User struct {
	ID       int    `form : id json : id`
	Name     string `form : name json : name`
	Email    string `form : email json : email`
	Password string `form : password json : password`
	ListFilm []Film `form : listfilm json : listfilm`
}

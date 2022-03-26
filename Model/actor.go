package Model

type Actor struct {
	ID        int    `form : id json : id`
	Name      string `form : name json : name`
	BirthDate string `form : birthdate json : birthdate`
	Age       int    `form : age json : age`
	Gender    string `form : gender json : gender`
}

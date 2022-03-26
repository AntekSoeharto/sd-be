package Model

type DetailActor struct {
	ID     int    `form : id json : id`
	FilmID string `form : filmid json : filmid`
	Actor  Actor  `form : id json : id`
}

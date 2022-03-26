package Model

type Film struct {
	ID            int       `form : id json : id`
	Judul         string    `form : judul json : judul`
	Rating        float64   `form : rating json : rating`
	TanggalTerbit string    `form : tanggalterbit json : tanggalterbit`
	ListComment   []Comment `form : listcomment json : listcomment`
	Actor         string    `form : listactor json : listactor`
	Sinopsis      string    `form : sinopsis json : sinopsis`
	FilmType      string    `form : filmtype json : filmtype`
	ReleaseType   string    `form : releasetype json : releasetype`
	Duration      int       `form : duration json : duration`
}

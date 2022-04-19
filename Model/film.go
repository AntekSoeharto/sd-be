package Model

type Film struct {
	ID              int       `form : id json : id`
	Judul           string    `form : judul json : judul`
	Rating          float64   `form : rating json : rating`
	TanggalTerbit   string    `form : tanggalterbit json : tanggalterbit`
	Actor           string    `form : listactor json : listactor`
	Sinopsis        string    `form : sinopsis json : sinopsis`
	Genre           string    `form : gender json : gender`
	FilmType        string    `form : filmtype json : filmtype`
	ReleaseType     string    `form : releasetype json : releasetype`
	Duration        int       `form : duration json : duration`
	DisplayMethod   int       `form : display_method json : display_method`
	Image           string    `form : image json : image`
	ImageBackground string    `form : image_background json : image_background`
}

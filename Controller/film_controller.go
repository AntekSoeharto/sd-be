package Controller

import (
	"fmt"
	"log"
	"net/http"
	"sd-api/Model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddFilm(c *gin.Context) {
	db := connect()
	defer db.Close()

	judul := c.PostForm("judul")
	rating := c.PostForm("rating")
	tanggal_terbit := c.PostForm("tanggal_terbit")
	actor := c.PostForm("actor")
	sinopsis := c.PostForm("sinopsis")
	genre := c.PostForm("genre")
	film_type := c.PostForm("film_type")
	release_type := c.PostForm("release_type")
	duration, _ := strconv.Atoi(c.PostForm("duration"))
	image := c.PostForm("image")
	img_background := c.PostForm("img_background")
	fmt.Println(judul)
	fmt.Println(rating)
	fmt.Println(tanggal_terbit)
	fmt.Println(actor)
	fmt.Println(sinopsis)
	fmt.Println(genre)
	fmt.Println(film_type)
	fmt.Println(release_type)
	fmt.Println(duration)

	_, errQuery := db.Exec("INSERT INTO films(judul, rating, tanggal_terbit, actor, sinopsis, genre, film_type, release_type, duration, display_method, image, img_background) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)",
		judul,
		rating,
		tanggal_terbit,
		actor,
		sinopsis,
		genre,
		film_type,
		release_type,
		duration,
		0,
		image,
		img_background,
	)
	print(errQuery)

	var response Model.ResponseData
	if errQuery == nil {
		response.Status = 200
		response.Message = "Insert Film Success"
	} else {
		response.Status = 400
		response.Message = "Insert Film Failed!\n" + errQuery.Error()
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func RepoGetAllFilm(id string, film_type string, film_display string) []Model.Film {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM films"

	if id != "" {
		query += " WHERE id = " + id
	}

	if film_type != "" {
		query += " WHERE film_type = '" + film_type + "' AND display_method = " + film_display
	}

	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	var Film Model.Film
	var Films []Model.Film
	for rows.Next() {
		if err := rows.Scan(&Film.ID, &Film.Judul, &Film.Rating, &Film.TanggalTerbit,
			&Film.Actor, &Film.Sinopsis, &Film.Genre, &Film.FilmType, &Film.ReleaseType,
			&Film.Duration, &Film.DisplayMethod, &Film.Image, &Film.ImageBackground); err != nil {
			log.Fatal(err.Error())
		} else {
			Films = append(Films, Film)
		}
	}

	return Films
}

func GetAllFilm(c *gin.Context) {
	film_id := c.PostForm("judul")
	film_type := c.Query("film_type")
	film_display := c.Query("film_display")
	fmt.Println(film_type, film_display, film_display)

	Films := RepoGetAllFilm(film_id, film_type, film_display)

	var response Model.ResponseData
	if len(Films) > 0 {
		response.Status = 200
		response.Message = "Get Success"
		response.Data = Films
	} else {
		response.Status = 400
		response.Message = "Get Failed!"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func UpdateFilm(c *gin.Context) {
	db := connect()
	defer db.Close()

	judul := c.PostForm("judul")
	rating, _ := strconv.ParseFloat(c.PostForm("rating"), 32)
	tanggal_terbit := c.PostForm("tanggal_terbit")
	actor := c.PostForm("actor")
	sinopsis := c.PostForm("sinopsis")
	film_type := c.PostForm("film_type")
	release_type := c.PostForm("release_type")
	duration, _ := strconv.Atoi(c.PostForm("duration"))
	film_id := c.PostForm("film_id")

	var Film Model.Film = RepoGetAllFilm(film_id, "", "")[0]

	if judul != "" {
		Film.Judul = judul
	}
	if rating != 0 {
		Film.Rating = rating
	}
	if actor != "" {
		Film.Actor = actor
	}
	if sinopsis != "" {
		Film.Sinopsis = sinopsis
	}
	if tanggal_terbit != "" {
		Film.TanggalTerbit = tanggal_terbit
	}
	if film_type != "" {
		Film.FilmType = film_type
	}
	if release_type != "" {
		Film.ReleaseType = release_type
	}
	if duration != 0 {
		Film.Duration = duration
	}

	_, errQuery := db.Exec("UPDATE films SET judul=?, rating=?, tanggal_terbit=?, actor=?, sinopsis=?, film_type=?, release_type=?, duration=? WHERE id=?",
		Film.Judul,
		Film.Rating,
		Film.TanggalTerbit,
		Film.Actor,
		Film.Sinopsis,
		Film.FilmType,
		Film.ReleaseType,
		Film.Duration,
		Film.ID,
	)

	var response Model.ResponseData
	if errQuery == nil {
		response.Status = 200
		response.Message = "Update Success"
	} else {
		response.Status = 400
		response.Message = "Update Failed!"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func DeleteFilm(c *gin.Context) {
	db := connect()
	defer db.Close()

	film_id := c.Param("film_id")

	_, errQuery := db.Exec("DELETE FROM films WHERE id = ?",
		film_id,
	)

	var response Model.ResponseData
	if errQuery == nil {
		response.Status = 200
		response.Message = "Delete Success"
	} else {
		response.Status = 400
		response.Message = "Delete Failed!"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func SearchFilm(c *gin.Context) {
	db := connect()
	defer db.Close()

	search := c.Query("search")
	fmt.Println(search)

	query := "SELECT * FROM films WHERE judul LIKE '%" + search + "%'"

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	var Film Model.Film
	var Films []Model.Film
	for rows.Next() {
		if err := rows.Scan(&Film.ID, &Film.Judul, &Film.Rating, &Film.TanggalTerbit,
			&Film.Actor, &Film.Sinopsis, &Film.Genre, &Film.FilmType, &Film.ReleaseType,
			&Film.Duration, &Film.DisplayMethod, &Film.Image, &Film.ImageBackground); err != nil {
			log.Fatal(err.Error())
		} else {
			Films = append(Films, Film)
		}
	}

	var response Model.ResponseData
	if len(Films) > 0 {
		response.Status = 200
		response.Message = "Search Success"
		response.Data = Films
	} else {
		response.Status = 400
		response.Message = "Search Failed!"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

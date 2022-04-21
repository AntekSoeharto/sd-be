package Controller

import (
	"log"
	"net/http"
	"sd-api/Model"
	
	"fmt"

	"github.com/gin-gonic/gin"
)

func AddMyList(c *gin.Context) {
	db := connect()
	defer db.Close()

	user_id := c.PostForm("user_id")
	film_id := c.PostForm("film_id")

	_, errQuery := db.Exec("INSERT INTO my_list(user_id, film_id) VALUES (?,?)",
		user_id,
		film_id,
	)
	
	fmt.Println("Debug, masuk ke line 25")

	var response Model.ResponseData
	if errQuery == nil {
		response.Status = 200
		response.Message = "Add MyList Success"
	} else {
		response.Status = 400
		response.Message = "Add MyList Film Failed!/n" + errQuery.Error()
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func RepoGetMyList(id string) []Model.Film {
	db := connect()
	defer db.Close()

	query := "SELECT b.* FROM my_list a JOIN films b ON a.film_id = b.id WHERE a.user_id = " + id
	

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}
	
	fmt.Println("Debug, masuk ke line 52")
	fmt.Println(err)

	var Film Model.Film
	var Films []Model.Film
	var image string
	for rows.Next() {
		if err := rows.Scan(&Film.ID, &Film.Judul, &Film.Rating, &Film.TanggalTerbit,
			&Film.Actor, &Film.Sinopsis, &Film.FilmType, &Film.ReleaseType,
			&Film.Duration, &image); err != nil {
			log.Fatal(err.Error())
		} else {
			Films = append(Films, Film)
		}
	}

	return Films
}

func DeleteMyList(c *gin.Context) {
	db := connect()
	defer db.Close()

	my_list_id := c.Param("my_list_id")

	_, errQuery := db.Exec("DELETE FROM my_list WHERE id = ?",
		my_list_id,
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

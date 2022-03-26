package Controller

import (
	"log"
	"net/http"
	"sd-api/Model"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	db := connect()
	defer db.Close()

	user_id := c.PostForm("user_id")
	film_id := c.PostForm("film_id")
	comment := c.PostForm("comment")

	_, errQuery := db.Exec("INSERT INTO comments(user_id, film_id, comment) VALUES (?,?,?)",
		user_id,
		film_id,
		comment,
	)

	var response Model.ResponseData
	if errQuery == nil {
		response.Status = 200
		response.Message = "Comment Success"
	} else {
		response.Status = 400
		response.Message = "Comment Film Failed!/n" + errQuery.Error()
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func RepoGetAllComment(id string) []Model.Comment {
	db := connect()
	defer db.Close()

	query := "SELECT a.id, a.comment, b.id, b.nama, b.email b FROM comments a JOIN users b ON a.user_id = b.id"

	if id != "" {
		query += " WHERE a.film_id = " + id
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	var Comment Model.Comment
	var Comments []Model.Comment
	for rows.Next() {
		if err := rows.Scan(&Comment.ID, &Comment.Comment, &Comment.User.ID, &Comment.User.Name, &Comment.User.Email); err != nil {
			log.Fatal(err.Error())
		} else {
			Comments = append(Comments, Comment)
		}
	}

	return Comments
}

func DeleteComment(c *gin.Context) {
	db := connect()
	defer db.Close()

	comment_id := c.Param("comment_id")

	_, errQuery := db.Exec("DELETE FROM comments WHERE id = ?",
		comment_id,
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

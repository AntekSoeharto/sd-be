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
	// user_id_temp := c.PostForm("user_id")
	// fmt.Println(user_id_temp)

	_, errQuery := db.Exec("INSERT INTO comments(user_id, film_id, comment) VALUES (?,?,?)",
		user_id,
		film_id,
		comment,
	)
	
	fmt.Println("Debug, masuk ke line 27")
	
	fmt.Println("Debug, masuk ke line 29")

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

func GetAllComment(c *gin.Context) {
	db := connect()
	defer db.Close()

	film_id := c.Query("film_id")

	query := "SELECT a.id, a.comment, b.username, b.email b FROM comments a JOIN users b ON a.user_id = b.id"

	if film_id != "" {
		query += " WHERE a.film_id = " + film_id
	}
	
	fmt.Println("Debug, masuk ke line 54")

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	var Comment Model.Comment
	var Comments []Model.Comment
	for rows.Next() {
		if err := rows.Scan(&Comment.ID, &Comment.Comment, &Comment.Username, &Comment.Email); err != nil {
			log.Fatal(err.Error())
		} else {
			Comments = append(Comments, Comment)
		}
	}

	var response Model.ResponseData
	if len(Comments) > 0 {
		response.Status = 200
		response.Message = "Get Comment Success"
		response.Data = Comments
	} else {
		response.Status = 400
		response.Message = "Get Comment Failed!"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)

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

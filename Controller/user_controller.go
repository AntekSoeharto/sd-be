package Controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"sd-api/Model"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	nama := c.PostForm("nama")
	email := c.PostForm("email")
	password := GetMD5Hash(c.PostForm("password"))
	fmt.Println(nama)

	Users := RepoGetAllUser("")

	var sama bool = false

	for i := 0; i < len(Users); i++ {
		if Users[i].Email == email {
			sama = true
		}
	}
	fmt.Println(nama)

	var response Model.ResponseData
	if sama == false {
		_, errQuery := db.Exec("INSERT INTO users(nama, email, password) VALUES (?,?,?)",
			nama,
			email,
			password,
		)

		if errQuery == nil {
			response.Status = 200
			response.Message = "Insert Success"
		} else {
			response.Status = 400
			response.Message = "Insert Failed!\n" + errQuery.Error()
		}
	} else {
		response.Status = 400
		response.Message = "Insert Failed!\nEmail Sudah Terpakai"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func GetAllUser(c *gin.Context) {
	user_id := c.Param("user_id")

	Users := RepoGetAllUser(user_id)

	var response Model.ResponseData
	if len(Users) > 0 {
		response.Status = 200
		response.Message = "Get Success"
		response.Data = Users
	} else {
		response.Status = 400
		response.Message = "Get Failed!"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func RepoGetAllUser(id string) []Model.User {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM users"

	if id != "" {
		query += " WHERE id = " + id
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}
	fmt.Println("Masuk")

	var User Model.User
	var Users []Model.User
	for rows.Next() {
		if err := rows.Scan(&User.ID, &User.Name, &User.Email, &User.Password); err != nil {
			log.Fatal(err.Error())
		} else {
			Users = append(Users, User)
		}
	}

	return Users
}

func UpdateUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	nama := c.PostForm("nama")
	email := c.PostForm("email")
	password := c.PostForm("password")
	user_id := c.PostForm("user_id")

	var User Model.User = RepoGetAllUser(user_id)[0]

	if nama != "" {
		User.Name = nama
	}
	if email != "" {
		User.Email = email
	}
	if password != "" {
		User.Password = password
	}

	_, errQuery := db.Exec("UPDATE users SET nama=?, email=?, password=? WHERE id=?",
		User.Name,
		User.Email,
		User.Password,
		User.ID,
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

func DeleteUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	user_id := c.Param("user_id")

	_, errQuery := db.Exec("DELETE FROM users WHERE id = ?",
		user_id,
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

func Login(c *gin.Context) {
	db := connect()
	defer db.Close()

	email := c.Query("email")
	password := GetMD5Hash(c.Query("password"))
	var success bool = false

	Users := RepoGetAllUser("")
	for i := 0; i < len(Users); i++ {
		if Users[i].Email == email {
			if Users[i].Password == password {
				success = true
			}
		}
	}

	var response Model.ResponseData
	if success == true {
		response.Status = 200
		response.Message = "Login Success"
	} else {
		response.Status = 400
		response.Message = "Login Failed!"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

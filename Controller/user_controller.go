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

	Users := RepoGetAllUser("", "")

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
	user_id := c.Param("id")
	email := c.Query("email")

	Users := RepoGetAllUser(user_id, email)

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

func RepoGetAllUser(id string, email string) []Model.User {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM users"

	if id != "" {
		query += " WHERE id = " + id
	}

	if email != "" {
		query += " WHERE email = '" + email + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}
	fmt.Println("Masuk")

	var User Model.User
	var Users []Model.User
	for rows.Next() {
		if err := rows.Scan(&User.ID, &User.Name, &User.Username, &User.Email, &User.Gender, &User.BirthDay, &User.Password); err != nil {
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

	email := c.Query("email")
	password := c.Query("password")
	passwordHashMD5 := GetMD5Hash(password)

	_, errQuery := db.Exec("UPDATE users SET password=? WHERE email=?",
		passwordHashMD5,
		email,
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
	var User Model.User

	Users := RepoGetAllUser("", "")
	for i := 0; i < len(Users); i++ {
		if Users[i].Email == email {
			if Users[i].Password == password {
				User = Users[i]
				success = true
			}
		}
	}

	var response Model.ResponseData
	if success == true {
		response.Status = 200
		response.Message = "Login Success"
		response.Data = User
	} else {
		response.Status = 400
		response.Message = "Login Failed!"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func SignUp(c *gin.Context) {
	db := connect()
	defer db.Close()

	nama := c.Query("nama")
	username := c.Query("username")
	email := c.Query("email")
	gender := c.Query("gender")
	birthday := c.Query("birthday")
	password := c.Query("password")
	passwordHash := GetMD5Hash(password)
	success := true

	Users := RepoGetAllUser("", "")
	for i := 0; i < len(Users); i++ {
		if Users[i].Email == email {
			success = false
		}
	}

	if success == true {
		_, errQuery := db.Exec("INSERT INTO users(nama, username, email, gender, birthday, password) VALUES (?,?,?,?,?,?)",
			nama,
			username,
			email,
			gender,
			birthday,
			passwordHash,
		)
		if errQuery != nil {
			fmt.Println(errQuery)
			success = false
		}
	}

	var response Model.ResponseData
	if success == true {
		response.Status = 200
		response.Message = "SignUp Success"
	} else {
		response.Status = 400
		response.Message = "SignUp Failed!"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)

}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

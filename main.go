package main

import (
	"fmt"

	controller "sd-api/Controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()

	//FILM
	router.POST("/film", controller.AddFilm)
	router.GET("/film", controller.GetAllFilm)
	router.GET("/search", controller.SearchFilm)
	router.GET("/film/:film_id", controller.GetAllFilm)
	router.PUT("/film", controller.UpdateFilm)
	router.DELETE("/film/:film_id", controller.DeleteFilm)

	//User
	router.POST("/user", controller.AddUser)
	router.GET("/user", controller.GetAllUser)
	router.GET("/user/:user_id", controller.GetAllUser)
	router.PUT("/user", controller.UpdateUser)
	router.DELETE("/user/:user_id", controller.DeleteUser)
	router.GET("/login", controller.Login)
	router.POST("/signup", controller.SignUp)

	//Comments
	router.POST("/comment", controller.AddComment)
	router.GET("/comment", controller.GetAllComment)
	router.DELETE("/comment/:comment_id", controller.DeleteComment)

	//MyList
	router.POST("/my_list", controller.AddMyList)
	router.DELETE("/my_list/:my_list_id", controller.DeleteMyList)

	// router := mux.New/Router()

	// //Actor
	// router.HandleFunc("/actor", controller.InsertActor).Methods("POST")
	// router.HandleFunc("/actor", controller.GetAllActor).Methods("GET")
	// router.HandleFunc("/actor/{actor_id}", controller.GetActor).Methods("GET")
	// router.HandleFunc("/actor/{actor_id}", controller.UpdateActor).Methods("PUT")
	// router.HandleFunc("/actor/{actor_id}", controller.DeleteActor).Methods("DELETE")

	// //Film
	// router.HandleFunc("/film", controller.AddFilm).Methods("POST")

	// http.Handle("/", router)
	// fmt.Println("Connected to port 9090")
	// log.Fatal(http.ListenAndServe(":9090", router))
	router.Run(":9090")
	fmt.Println("Connected to port 9090")
}

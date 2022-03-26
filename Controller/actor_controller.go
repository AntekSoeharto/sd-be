package Controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sd-api/Model"
	"strconv"

	"github.com/gorilla/mux"
)

func InsertActor(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	Name := r.Form.Get("name")
	BirthDate := r.Form.Get("birth_date")
	Age, _ := strconv.Atoi(r.Form.Get("age"))
	Gender := r.Form.Get("gender")

	_, errQuery := db.Exec("INSERT INTO actors(name, birth_date, age, gender) VALUES (?, ?, ?, ?)",
		Name,
		BirthDate,
		Age,
		Gender,
	)

	var response Model.ResponseData
	if errQuery == nil {
		response.Status = 200
		response.Message = "Insert Actor Success"
	} else {
		response.Status = 400
		response.Message = "Insert Actor Failed!\n" + errQuery.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetAllActor(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM actors"

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	var Actor Model.Actor
	var Actors []Model.Actor
	for rows.Next() {
		if err := rows.Scan(&Actor.ID, &Actor.Name, &Actor.BirthDate, &Actor.Age,
			&Actor.Gender); err != nil {
			log.Fatal(err.Error())
		} else {
			Actors = append(Actors, Actor)
		}
	}

	var response Model.ResponseData
	if len(Actors) > 0 {
		response.Status = 200
		response.Message = "Get Success"
		response.Data = Actors
	} else {
		response.Status = 400
		response.Message = "Get Failed!"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetActor(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	query := "SELECT * FROM actors where id = " + id

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	var Actor Model.Actor
	for rows.Next() {
		if err := rows.Scan(&Actor.ID, &Actor.Name, &Actor.BirthDate, &Actor.Age,
			&Actor.Gender); err != nil {
			log.Fatal(err.Error())
		}
	}

	var response Model.ResponseData
	if err == nil {
		response.Status = 200
		response.Message = "Get Success"
		response.Data = Actor
	} else {
		response.Status = 400
		response.Message = "Get Failed!\n" + err.Error()
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetActorForUpdate(id string) Model.Actor {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM actors where id = " + id

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	var Actor Model.Actor
	for rows.Next() {
		if err := rows.Scan(&Actor.ID, &Actor.Name, &Actor.BirthDate, &Actor.Age,
			&Actor.Gender); err != nil {
			log.Fatal(err.Error())
		}
	}

	return Actor
}

func UpdateActor(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	Name := r.Form.Get("name")
	BirthDate := r.Form.Get("birth_date")
	Age, _ := strconv.Atoi(r.Form.Get("age"))
	Gender := r.Form.Get("gender")

	vars := mux.Vars(r)
	ActorId := vars["actor_id"]

	var ActorForUpdate Model.Actor = GetActorForUpdate(ActorId)

	if Name == "" {
		Name = ActorForUpdate.Name
	}
	if BirthDate == "" {
		BirthDate = ActorForUpdate.BirthDate
	}
	if Age == 0 {
		Age = ActorForUpdate.Age
	}
	if Gender == "" {
		Gender = ActorForUpdate.Gender
	}

	_, errQuery := db.Exec("UPDATE actors SET name=?, birth_date=?, age=?, gender=? WHERE id=?",
		Name,
		BirthDate,
		Age,
		Gender,
		ActorId,
	)

	var response Model.ResponseData
	if errQuery == nil {
		response.Status = 200
		response.Message = "Update Actor Success"
	} else {
		response.Status = 400
		response.Message = "Update Actor Failed!"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteActor(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	ActorId := vars["actor_id"]

	_, errQuery := db.Exec("DELETE FROM actors WHERE id = ?",
		ActorId,
	)

	var response Model.ResponseData
	if errQuery == nil {
		response.Status = 200
		response.Message = "Delete Actor Success"
	} else {
		response.Status = 400
		response.Message = "Delete Actor Failed!"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

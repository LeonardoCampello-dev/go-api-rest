package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LeonardoCampello-dev/go-api-rest/db"
	"github.com/LeonardoCampello-dev/go-api-rest/models"
	"github.com/gorilla/mux"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Home Page")
}

func GetAllPersonalities(writer http.ResponseWriter, request *http.Request) {
	var personalities []models.Personality

	db.DB.Find(&personalities)

	json.NewEncoder(writer).Encode(personalities)
}

func GetPersonalityById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	var personality models.Personality

	db.DB.First(&personality, id)

	json.NewEncoder(writer).Encode(personality)
}

func CreatePersonality(writer http.ResponseWriter, request *http.Request) {
	var personality models.Personality

	json.NewDecoder(request.Body).Decode(&personality)

	db.DB.Create(&personality)

	json.NewEncoder(writer).Encode(personality)
}

func DeletePersonality(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	var personality models.Personality

	db.DB.Delete(&personality, id)
}

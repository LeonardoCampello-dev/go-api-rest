package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeonardoCampello-dev/go-api-rest/models"
	"github.com/gorilla/mux"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Home Page")
}

func GetAllPersonalities(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(models.Personalities)
}

func GetPersonalityById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	for _, personality := range models.Personalities {
		if strconv.Itoa(personality.Id) == id {
			json.NewEncoder(writer).Encode(personality)
		}
	}
}

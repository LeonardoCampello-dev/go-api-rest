package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LeonardoCampello-dev/go-api-rest/models"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Home Page")
}

func GetAllPersonalities(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(models.Personalities)
}

package routes

import (
	"log"
	"net/http"

	"github.com/LeonardoCampello-dev/go-api-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()

	http.Handle("/", router)

	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/personalities", controllers.GetAllPersonalities).Methods("GET")
	router.HandleFunc("/personalities/{id}", controllers.GetPersonalityById).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", nil))
}

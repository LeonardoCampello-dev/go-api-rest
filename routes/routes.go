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
	router.HandleFunc("/personalities", controllers.GetAllPersonalities)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

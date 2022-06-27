package routes

import (
	"log"
	"net/http"

	"github.com/LeonardoCampello-dev/go-api-rest/controllers"
	"github.com/LeonardoCampello-dev/go-api-rest/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()

	router.Use(middlewares.ContentTypeMiddleware)

	http.Handle("/", router)

	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/personalities", controllers.GetAllPersonalities).Methods("GET")
	router.HandleFunc("/personalities/{id}", controllers.GetPersonalityById).Methods("GET")
	router.HandleFunc("/personalities", controllers.CreatePersonality).Methods("POST")
	router.HandleFunc("/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	router.HandleFunc("/personalities/{id}", controllers.UpdatePersonality).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}

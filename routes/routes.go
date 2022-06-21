package routes

import (
	"log"
	"net/http"

	"github.com/LeonardoCampello-dev/go-api-rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/personalities", controllers.GetAllPersonalities)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

package main

import (
	"fmt"

	"github.com/LeonardoCampello-dev/go-api-rest/db"
	"github.com/LeonardoCampello-dev/go-api-rest/routes"
)

func main() {
	db.ConnectDB()

	fmt.Println("Go API Rest")

	routes.HandleRequest()
}

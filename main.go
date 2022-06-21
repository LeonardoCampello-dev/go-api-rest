package main

import (
	"fmt"

	"github.com/LeonardoCampello-dev/go-api-rest/routes"
)

func main() {
	fmt.Println("Go API Rest")

	routes.HandleRequest()
}

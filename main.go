package main

import (
	"fmt"

	"github.com/LeonardoCampello-dev/go-api-rest/models"
	"github.com/LeonardoCampello-dev/go-api-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Test 1", History: "Test 1"},
		{Name: "Test 2", History: "Test 2"},
	}

	fmt.Println("Go API Rest")

	routes.HandleRequest()
}

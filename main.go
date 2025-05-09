package main

import (
	"net/http"

	"github.com/alvinabner/crud-employee-go/database"
	"github.com/alvinabner/crud-employee-go/routes"
)

func main() {
	database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server)

	http.ListenAndServe(":8080", server)
}

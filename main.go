package main

import (
	"net/http"

	"github.com/alvinabner/crud-employee-go/controller"
	"github.com/alvinabner/crud-employee-go/database"
)

func main() {
	database.InitDatabase()

	server := http.NewServeMux()

	server.HandleFunc("/", controller.NewHelloWorldController())

	http.ListenAndServe(":8080", server)
}

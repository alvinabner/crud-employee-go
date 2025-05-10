package routes

import (
	"database/sql"
	"net/http"

	"github.com/alvinabner/crud-employee-go/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.NewHelloWorldController())
	server.HandleFunc("/employee", controller.NewIndexEmployee())
	server.HandleFunc("/employee/create", controller.NewCreateEmployeeController(db))
}

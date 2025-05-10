package main

import (
	"fmt"
	"net/http"

	"github.com/alvinabner/crud-employee-go/database"
	"github.com/alvinabner/crud-employee-go/routes"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	fmt.Println("Server sedang berjalan di http://localhost:8080")
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}

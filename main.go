package main

import (
	"net/http"

	"github.com/alvinabner/crud-employee-go/database"
)

func main() {
	database.InitDatabase()

	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":8080", server)
	// fmt.Println("Hello World")
}

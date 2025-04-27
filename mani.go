package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/balamgit/DAGaroo/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/about", handlers.AboutHandler)

	fmt.Println("Server is running at http://localhost:8081")
	http.ListenAndServe(":8081", r)
}

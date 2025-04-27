package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/balamgit/DAGaroo/internal/http/controllers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.HomeHandler)
	r.HandleFunc("/about", controllers.AboutHandler)

	fmt.Println("Server is running at http://localhost:8081")
	http.ListenAndServe(":8081", r)
}

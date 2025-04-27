package controllers

import (
	"fmt"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "DAGaroo about page.")
}

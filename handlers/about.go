package handlers

import (
	"fmt"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ℹ️ This is the About Page.")
}

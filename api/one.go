package api

import (
	"fmt"
	"net/http"
)

func OneHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Hello from one!")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

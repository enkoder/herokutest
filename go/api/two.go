package api

import (
	"fmt"
	"net/http"
)

func TwoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Hello from two!")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

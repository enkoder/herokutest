package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-web/api"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	// Using gorilla for routing and enabling cors
	router := mux.NewRouter()
	router.HandleFunc("/", api.OneHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

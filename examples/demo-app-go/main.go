package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
    port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

    r := mux.NewRouter()
    r.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World from a Go App!")
	}))

	log.Fatal(http.ListenAndServe(":"+port, r))
}

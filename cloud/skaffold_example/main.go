package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, I'm skaffold!")
}

func main() {
	log.Println("Skaffold server") // change logging to see the differences
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

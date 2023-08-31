package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %q Server\n", html.EscapeString(r.Proto))
}

func main() {
	http.HandleFunc("/", getRoot)
	fmt.Println("Server running at 127.0.0.1:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

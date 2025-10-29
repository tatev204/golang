package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

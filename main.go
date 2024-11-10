package main

import (
	"fmt"
	"net/http"
)

func HandleHelloWorld(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	fmt.Println("Program starting...")

	http.HandleFunc("/", HandleHelloWorld)

	http.ListenAndServe(":8081", nil)
}

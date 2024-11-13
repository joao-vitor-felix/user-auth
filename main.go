package main

import (
	"fmt"
	"net/http"

	"github.com/goschool/crud/routes"
)

func main() {
	r := routes.SetupRoutes()
	fmt.Println("Program starting...")
	http.ListenAndServe(":8081", r)
}

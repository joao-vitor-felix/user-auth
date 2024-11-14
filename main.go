package main

import (
	"fmt"
	"net/http"

	"github.com/goschool/crud/db"
	"github.com/goschool/crud/routes"
)

func main() {
	r := routes.SetupRoutes()
	_, err := db.Open()

	if err != nil {
		panic(err)
	}

	fmt.Println("Program starting...")
	http.ListenAndServe(":8081", r)
}

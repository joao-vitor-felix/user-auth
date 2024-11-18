package main

import (
	"fmt"
	"net/http"

	"github.com/goschool/crud/api"
	"github.com/goschool/crud/db"
	"github.com/goschool/crud/routes"
)

func main() {
	database, err := db.Open()
	if err != nil {
		panic(err)
	}

	userStore := db.NewSQLiteUserStore(database)
	userHandler := api.NewUserHandler(userStore)
	r := routes.SetupRoutes(*userHandler)

	fmt.Println("Program starting...")
	http.ListenAndServe(":8081", r)
}

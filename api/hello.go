package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goschool/crud/types"
)

func HandleHelloWorld(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func HandleEchoUser(w http.ResponseWriter, r *http.Request) {
	var user types.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "this is the email: %s", user.Email)
}

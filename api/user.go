package api

import (
	"encoding/json"
	"net/http"

	"github.com/goschool/crud/db"
	"github.com/goschool/crud/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(us db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: us,
	}
}

func (u *UserHandler) HandlerRegisterUser(w http.ResponseWriter, r *http.Request) {
	var createUser types.CreateUser
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&createUser); err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
	}

	user, err := types.NewUser(createUser)
	if err != nil {
		http.Error(w, "Invalid new user", http.StatusBadRequest)
		return
	}

	newUser, err := u.userStore.CreateUser(ctx, user)
	if err != nil {
		http.Error(w, "Invalid new user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		http.Error(w, "Could not register new user", http.StatusInternalServerError)
	}
}

func (u *UserHandler) HandlerLoginUser(w http.ResponseWriter, r *http.Request) {

}

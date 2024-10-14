package api

import (
	"net/http"

	"github.com/goschool/crud/db"
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
}

func (u *UserHandler) HandlerLoginUser(w http.ResponseWriter, r *http.Request) {

}

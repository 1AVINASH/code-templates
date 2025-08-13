package user

import (
	"net/http"
)

type UserManager struct {
	repo *UserRepository
	apis *UserAPIs
}

func NewUserManager(mux *http.ServeMux) *UserManager {
	repo := NewUserRepository()
	apis := NewUserApis(mux, repo)
	return &UserManager{
		repo: repo,
		apis: apis,
	}
}

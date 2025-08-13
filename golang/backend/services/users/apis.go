package user

import "net/http"

type UserAPIs struct {
	repo *UserRepository
}

func NewUserApis(mux *http.ServeMux, repo *UserRepository) *UserAPIs {
	ua := &UserAPIs{repo: repo}
	ua.registerUserRoutes(mux)
	return ua
}

func (ua *UserAPIs) registerUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users", ua.getUsers)
    mux.HandleFunc("/users/create", ua.createUser)
}


func (ua *UserAPIs) getUsers(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("List of users"))
}

func (ua *UserAPIs) createUser(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("User created"))
}
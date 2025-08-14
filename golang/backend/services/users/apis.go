package user

import (
	"gotemplate/utility/logger"
	middleware "gotemplate/utility/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserAPIs struct {
	repo *UserRepository
}

func NewUserApis(mux *chi.Mux, repo *UserRepository) *UserAPIs {
	ua := &UserAPIs{repo: repo}
	ua.registerUserRoutes(mux)
	return ua
}

func (ua *UserAPIs) registerUserRoutes(mux *chi.Mux) {
	mux.Route("/users", func(mux chi.Router) {
		mux.Get("/", middleware.JSON(ua.getUsers))
		mux.Post("/", middleware.JSON(ua.createUser))
	})
}

func (ua *UserAPIs) getUsers(w http.ResponseWriter, r *http.Request) (interface{}, int) {
	logger.GetLogger().Infof("Starting to get user")
	return []string{"Alice", "Bob"}, http.StatusOK
}

func (ua *UserAPIs) createUser(w http.ResponseWriter, r *http.Request) (interface{}, int) {
	logger.GetLogger().Infof("Starting to create user")
	return map[string]string{"status": "created"}, http.StatusCreated
}

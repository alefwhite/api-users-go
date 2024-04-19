package routes

import (
	"github.com/alefwhite/api-users-go/config/env"
	"github.com/alefwhite/api-users-go/internal/handler"
	"github.com/alefwhite/api-users-go/internal/handler/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

func InitUserRoutes(router chi.Router, h handler.UserHandler) {
	router.Use(middleware.LoggerData)

	router.Post("/user", h.CreateUser)
	router.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(env.Env.TokenAuth))
		r.Use(jwtauth.Authenticator)

		//user routes
		r.Patch("/user", h.UpdateUser)
		r.Get("/user", h.GetUserByID)
		r.Delete("/user", h.DeleteUser)
		r.Get("/user/list-all", h.FindManyUsers)
		r.Patch("/user/password", h.UpdateUserPassword)
	})
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
	})
}

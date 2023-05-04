package routes

import (
	"jwt/handlers"

	"github.com/go-chi/chi"
)

func Routes(router *chi.Mux) {
	// group other routes with /api
	router.Route("/api", func(r chi.Router) {
		UserRoutes(r)
		AuthRoutes(r)
	})
}

func UserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.UserListHandler)
		r.Get("/{id}", handlers.UserGetByIDHandler)
		r.Post("/", handlers.UserInsertHandler)
		r.Put("/{id}", handlers.UserUpdateHandler)
		r.Delete("/{id}", handlers.DeleteUserHandler)
	})
}

func AuthRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/otp", handlers.GetOTPHandler)
		r.Post("/login", handlers.LoginHandler)
	})
}

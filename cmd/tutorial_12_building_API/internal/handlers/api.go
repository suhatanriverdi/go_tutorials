package handlers

import (
	"github.com/avukadin/goapi/internal/middleware"
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
)

// "H" capital letter means public and can be imported by other packages.
// "h" lowercase letter means private and can only be used within the package.
func Handler(r *chi.Mux) {
	// Global middleware
	// Applied all the time
	// This middleware is used to strip slashes from the URL path
	// For example, ".../users/" will be converted to ".../users"
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// Middleware for account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}

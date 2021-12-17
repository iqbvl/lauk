package fetch

import (
	"time"

	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/rs/cors"
)

func (d *REST) RegisterRoute(r *chi.Mux) {
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)

	// Basic CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(cors.Handler)
	r.Use(chimiddleware.Timeout(60 * time.Second))

	//protected
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(d.TokenJWT))
		r.Use(jwtauth.Authenticator)
		r.Get("/get", d.GetStorageHandler)
		r.Get("/validatejwt", d.ValidateJWTHandler)
	})
}

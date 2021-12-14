package auth

import (
	"github.com/go-chi/chi"
)

type Delivery interface {
	RegisterRoute(r *chi.Mux)
}

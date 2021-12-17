package delivery

import (
	"github.com/go-chi/chi"
)

type Auth interface {
	RegisterRoute(r *chi.Mux)
}

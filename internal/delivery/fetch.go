package delivery

import (
	"github.com/go-chi/chi"
)

type Fetch interface {
	RegisterRoute(r *chi.Mux)
}

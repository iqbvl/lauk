package external

import (
	"net/http"

	"github.com/iqbvl/lauk/internal/repository"
)

type External struct {
	Client *http.Client
}

func NewExternal(e External) repository.External {
	return &External{
		Client: e.Client,
	}
}

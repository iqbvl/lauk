package fetch

import (
	"context"

	"github.com/go-chi/jwtauth"
	"github.com/iqbvl/lauk/internal/delivery"
	"github.com/iqbvl/lauk/internal/usecase"
)

// var TokenAuth *jwtauth.JWTAuth

const (
	invalidEmailFormat      = "You are inputting wrong email address format"
	emptyLoginField         = "Username and Password cant be empty"
	postMethodSupported     = "Only Post Allowed"
	getMethodSupported      = "Only GET Allowed"
	errorConvertRequestBody = "Error when converting request body"
	tokenError              = "Error Generate Token"
)

type REST struct {
	Context      context.Context
	FetchUsecase usecase.Fetch
	TokenJWT     *jwtauth.JWTAuth
}

func NewREST(args REST) delivery.Fetch {
	return &REST{
		FetchUsecase: args.FetchUsecase,
		TokenJWT:     args.TokenJWT,
		Context:      args.Context,
	}
}

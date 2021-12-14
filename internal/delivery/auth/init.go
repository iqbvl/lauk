package auth

import (
	"context"

	"github.com/go-chi/jwtauth"
	"github.com/iqbvl/lauk/internal/usecase"
)

// var TokenAuth *jwtauth.JWTAuth

const (
	invalidEmailFormat      = "You are inputting wrong email address format"
	emptyLoginField         = "Username and Password cant be empty"
	postMethodSupported     = "Only Post Allowed"
	errorConvertRequestBody = "Error when converting request body"
	tokenError              = "Error Generate Token"
)

type REST struct {
	Context     context.Context
	AuthUsecase usecase.Auth
	TokenJWT    *jwtauth.JWTAuth
}

func NewREST(args REST) Delivery {
	return &REST{
		AuthUsecase: args.AuthUsecase,
		TokenJWT:    args.TokenJWT,
		Context:     args.Context,
	}
}

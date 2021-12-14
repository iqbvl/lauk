package auth

import "github.com/iqbvl/lauk/internal/repository/auth"

type AuthUsecase struct {
	TTLCache auth.Cache
}

func NewUsecase(args AuthUsecase) *AuthUsecase{
	return&AuthUsecase{
		TTLCache: args.TTLCache,
	}
}
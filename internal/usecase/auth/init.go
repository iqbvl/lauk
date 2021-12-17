package auth

import "github.com/iqbvl/lauk/internal/repository"

type AuthUsecase struct {
	TTLCache repository.Cache
}

func NewUsecase(args AuthUsecase) *AuthUsecase{
	return&AuthUsecase{
		TTLCache: args.TTLCache,
	}
}
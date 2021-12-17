package fetch

import "github.com/iqbvl/lauk/internal/repository"

type FetchUsecase struct {
	External repository.External
	TTLCache repository.Cache
}

func NewUsecase(args FetchUsecase) *FetchUsecase {
	return &FetchUsecase{
		External: args.External,
		TTLCache: args.TTLCache,
	}
}

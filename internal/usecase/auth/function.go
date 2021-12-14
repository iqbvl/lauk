package auth

import (
	"context"

	"github.com/iqbvl/lauk/internal/model"
	"github.com/iqbvl/lauk/internal/platform/util"
	log "github.com/sirupsen/logrus"
)

func (u *AuthUsecase) GetUser(ctx context.Context, args model.User) (model.User, error) {
	user, err := u.TTLCache.GetUser(ctx, util.GenerateKey(args))
	if err != nil {
		if err.Error() != "key not found" {
			log.Errorf("[AuthUsecase][u.TTLCache.GetUser] msg : %s \n", err.Error())
			return user, err
		} 
	}
	return user, nil
}

func (u *AuthUsecase) GetUserExistence(ctx context.Context, args model.User) (bool, error) {
	user, err := u.TTLCache.GetUser(ctx, util.GenerateKey(args))
	if err != nil {
		log.Errorf("[AuthUsecase][u.TTLCache.GetUser] msg : %s \n", err.Error())
		return false, err
	}

	if user.Name != "" {
		return true, nil
	}

	log.Infof("[AuthUsecase] msg : %s \n", "user not exists")
	return false, nil
}

func (u *AuthUsecase) SetUser(ctx context.Context, args model.User) error {
	err := u.TTLCache.SetUser(ctx, args)
	if err != nil {
		log.Errorf("[AuthUsecase][u.TTLCache.SetUser] msg : %s \n", err.Error())
	}
	return nil
}

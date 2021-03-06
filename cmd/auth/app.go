package main

import (
	"context"
	"net/http"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/iqbvl/lauk/internal/delivery"
	"github.com/iqbvl/lauk/internal/delivery/auth"
	"github.com/iqbvl/lauk/internal/repository"
	rTTLCache "github.com/iqbvl/lauk/internal/repository/ttlcache"
	"github.com/iqbvl/lauk/internal/usecase"
	uAuth "github.com/iqbvl/lauk/internal/usecase/auth"
	log "github.com/sirupsen/logrus"
)

var (
	TokenAuth    *jwtauth.JWTAuth
	ttlCacheRepo repository.Cache
	authUsecase  usecase.Auth
	authDelivery delivery.Auth
	ctx          context.Context
)

func init() {
	//init repo
	ttlCacheRepo = rTTLCache.NewTTLCache(ttlcache.NewCache())

	//init usecase
	authUsecase = uAuth.NewUsecase(uAuth.AuthUsecase{
		TTLCache: ttlCacheRepo,
	})

	secretKey := "aXFiYWwgYWJkdXJyYWhtYW4="
	TokenAuth = jwtauth.New("HS256", []byte(secretKey), nil)

	//init delivery
	ctx = context.Background()
	authDelivery = auth.NewREST(auth.REST{Context: ctx, AuthUsecase: authUsecase, TokenJWT: TokenAuth})
}

func main() {
	router := chi.NewRouter()
	authDelivery.RegisterRoute(router)
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", router))
}

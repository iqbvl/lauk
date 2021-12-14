package main

import (
	"context"
	"net/http"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/iqbvl/lauk/internal/delivery/auth"
	rAuth "github.com/iqbvl/lauk/internal/repository/auth"
	rAuthTTLCache "github.com/iqbvl/lauk/internal/repository/auth/ttlcache"
	"github.com/iqbvl/lauk/internal/usecase"
	uAuth "github.com/iqbvl/lauk/internal/usecase/auth"
	log "github.com/sirupsen/logrus"
)

var (
	TokenAuth        *jwtauth.JWTAuth
	authTTLCacheRepo rAuth.Cache
	authUsecase      usecase.Auth
	authDelivery     auth.Delivery
	ctx              context.Context
)

func init() {
	authTTLCacheRepo = rAuthTTLCache.NewTTLCache(ttlcache.NewCache())
	authUsecase = uAuth.NewUsecase(uAuth.AuthUsecase{
		TTLCache: authTTLCacheRepo,
	})

	secretKey := "aXFiYWwgYWJkdXJyYWhtYW4="
	TokenAuth = jwtauth.New("HS256", []byte(secretKey), nil)

	ctx = context.Background()
	authDelivery = auth.NewREST(auth.REST{Context: ctx, AuthUsecase: authUsecase, TokenJWT: TokenAuth})

}

func main() {
	router := chi.NewRouter()
	authDelivery.RegisterRoute(router)
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", router))
}

package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/iqbvl/lauk/internal/delivery"
	"github.com/iqbvl/lauk/internal/delivery/fetch"
	"github.com/iqbvl/lauk/internal/repository"
	"github.com/iqbvl/lauk/internal/repository/external"
	rTTLCache "github.com/iqbvl/lauk/internal/repository/ttlcache"
	"github.com/iqbvl/lauk/internal/usecase"
	uFetch "github.com/iqbvl/lauk/internal/usecase/fetch"
)

var (
	TokenAuth     *jwtauth.JWTAuth
	externalRepo  repository.External
	ttlCacheRepo repository.Cache
	fetchUsecase  usecase.Fetch
	fetchDelivery delivery.Fetch
	ctx           context.Context
	httpClient    *http.Client
)

func init() {
	secretKey := "aXFiYWwgYWJkdXJyYWhtYW4="
	TokenAuth = jwtauth.New("HS256", []byte(secretKey), nil)

	ctx = context.Background()

	//http connection pooling
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	httpClient = &http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}

	//init repo
	externalRepo = external.NewExternal(external.External{Client: httpClient})
	ttlCacheRepo = rTTLCache.NewTTLCache(ttlcache.NewCache())
	
	//init usecase
	fetchUsecase = uFetch.NewUsecase(uFetch.FetchUsecase{
		External: externalRepo,
		TTLCache: ttlCacheRepo,
	})

	//init delivery
	fetchDelivery = fetch.NewREST(fetch.REST{Context: ctx, FetchUsecase: fetchUsecase, TokenJWT: TokenAuth})
}

func main() {
	router := chi.NewRouter()
	fetchDelivery.RegisterRoute(router)
	log.Fatal(http.ListenAndServe("0.0.0.0:6000", router))
}

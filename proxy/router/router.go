package router

import (
	"proxy/internal/controller"
	"proxy/middleware"

	"github.com/go-chi/chi"
)

type RouterConfig struct {
	router chi.Mux
}

type RouterOption func(*RouterConfig)

func NewRouter(options ...RouterOption) *RouterConfig {

	var router RouterConfig = RouterConfig{}

	for _, option := range options {
		option(&router)
	}

	return &router
}

func PublicRouterOption() RouterOption {
	return func(rc *RouterConfig) {
		var publicRouter *chi.Mux = chi.NewRouter()
		publicRouter.Post("/login", controller.HandleLogin)
		publicRouter.Post("/registration", nil) //добавь хендлер)
		rc.router.Mount("/public", publicRouter)
	}
}

func PrivateRouterOption() RouterOption {
	return func(rc *RouterConfig) {
		var protectedRouter *chi.Mux = chi.NewRouter()
		protectedRouter.Use(middleware.JWTAuthMiddleware)
		protectedRouter.Post("/address/geocode", controller.HandleGeoCode)
		protectedRouter.Post("/address/search", controller.SearchHandler)

		rc.router.Mount("/private", protectedRouter)
	}
}

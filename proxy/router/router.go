package router

import (
	"proxy/controller"
	"proxy/middleware"

	"github.com/go-chi/chi"
)

func SetupRouter() *chi.Mux {

	router := chi.NewRouter()

	//Публичные ссылки
	publicRouter := chi.NewRouter()
	publicRouter.Post("/login", controller.HandleLogin)
	publicRouter.Post("/registration", nil) //добавь хендлер
	router.Mount("/", publicRouter)

	//Приватные ссылки
	protectedRouter := chi.NewRouter()
	protectedRouter.Use(middleware.JWTAuthMiddleware)
	protectedRouter.Post("/address/geocode", controller.HandleGeoCode)
	protectedRouter.Post("/address/search", controller.HandleSearch)
	router.Mount("/api", protectedRouter)

	return router
}

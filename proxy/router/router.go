package router

import (
	"proxy/controller"
	"proxy/middleware"
	"proxy/service"

	"github.com/go-chi/chi"
)

func SetupRouter() *chi.Mux {
	searchService := service.NewSearch()                        // Создание сервиса
	searchHandler := controller.NewSearchHandler(searchService) // Создание хендлера

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
	protectedRouter.Post("/address/search", searchHandler)
	router.Mount("/api", protectedRouter)

	return router
}

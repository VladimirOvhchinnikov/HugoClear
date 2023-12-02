package service

import "errors"

// GeoResult представляет собой временную структуру для хранения результатов поиска.
type GeoResult struct {
	// как возможный пример формы ответа
	Result string
}

// GeoService управляет логикой геолокации.
type GeoService struct {
	// Структура контроллера геолокации, может быть расширена в будущем
}

// GeoServiceOption определяет тип функции опции для GeoController.
type GeoServiceOption func(*GeoService)

// NewGeoService создает новый экземпляр GeoController с применением переданных опций.
func NewGeoService(options ...GeoServiceOption) *GeoService {
	var controller GeoService = GeoService{}

	for _, option := range options {
		option(&controller)
	}

	return &controller
}

// GeoCoder определяет интерфейс для компонентов, выполняющих поиск.
type GeoCoder interface {
	GeoSE(query string) (GeoResult, error)
}

// GeoSE выполняет поиск и возвращает результаты.
func (g *GeoService) GeoSE(query string) (GeoResult, error) {

	var geoResult GeoResult = GeoResult{}

	if query != "" {
		geoResult = GeoResult{
			Result: "Все хорошо. Псевдо запрос обработан",
		}
	} else {
		return GeoResult{}, errors.New("query is empty")
	}

	return geoResult, nil
}

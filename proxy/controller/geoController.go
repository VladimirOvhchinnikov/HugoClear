package controller

import (
	"log"
	"mode/proxy/service"
	"mode/proxy/utils"
	"net/http"
)

// Рандомная структура для формы. потом переделай
type geoCodeUnm struct {
	Query string `json:"query"`
}

func (g geoCodeUnm) Process() error {
	return nil
}

func HandleGeoCode(w http.ResponseWriter, r *http.Request) {

	log.Println("HandleGeoCode - work now")

	// Определение формат прtшедшего запроса
	var extraData utils.ExtractDataFromRequest
	var err error
	extraData, err = extraData.Extract(r)
	if err != nil {
		log.Println(err)
		return
	}

	// АНмаршелинг тела запроса
	var geocode geoCodeUnm
	err = extraData.UnmarshalAndProcess(r, &geocode)
	if err != nil {
		log.Println(err)
	}

	//Отправка данных слою бизнес логики
	var geoLogic service.GeoService = *service.NewGeoService()
	var geoLogicResult service.GeoResult = service.GeoResult{}
	geoLogicResult, err = geoLogic.GeoSE(geocode.Query)

	//типо отправка данных. без подготовки данных
	// Кодирование и отправка ответа
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(geoLogicResult.Result))
}

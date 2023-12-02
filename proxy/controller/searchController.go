package controller

import (
	"log"
	"mode/proxy/service"
	"mode/proxy/utils"
	"net/http"
)

// Рандомная структура для формы. потом переделай
type searchUnm struct {
	Query string `json:"query"`
}

func (s searchUnm) Process() error {
	return nil
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {

	log.Println("HandleSearch - work now")

	//Определение формат прtшедшего запроса
	var extraData utils.ExtractDataFromRequest
	var err error
	extraData, err = extraData.Extract(r)
	if err != nil {
		log.Println(err)
		return
	}

	//АНмаршелинг тела запроса
	var search searchUnm
	err = extraData.UnmarshalAndProcess(r, &search)
	if err != nil {
		log.Println(err)
	}

	//Отправка данных слою бизнес логики
	var searchLogic service.SearchService = *service.NewSearchService()
	var searchLogicResult service.SearchResult = service.SearchResult{}
	searchLogicResult, err = searchLogic.Search(search.Query)

	//типо отправка данных. без подготовки данных
	// Кодирование и отправка ответа
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(searchLogicResult.Result))
}

package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"proxy/service"
)

// searchUnm is a temporary structure for the search form. Needs to be refined later.
type searchUnm struct {
	Query string `json:"query"`
}

func (s searchUnm) Process() error {
	return nil
}

// HandleSearch handles search requests.
// @Summary Search
// @Description This endpoint processes search requests and returns search results.
// @Tags Search
// @Accept  json
// @Produce  json
// @Param query body searchUnm true "Search Query"
// @Success 200 {object} string "Successfully processed search request"
// @Failure 400 {string} string "Bad Request"
// @Router /search [post]
// NewSearchHandler создаёт новый HTTP хендлер для поисковых запросов.
// Возвращает функцию, которая соответствует сигнатуре http.HandlerFunc.
func NewSearchHandler(searcher service.Searcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("HandleSearch - запрос обрабатывается")

		// Извлечение запроса из тела запроса
		var searchReq searchUnm
		if err := json.NewDecoder(r.Body).Decode(&searchReq); err != nil {
			http.Error(w, "Некорректный запрос", http.StatusBadRequest)
			return
		}

		// Передача данных в слой бизнес-логики
		result, err := searcher.Search(searchReq.Query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Кодирование и отправка ответа
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

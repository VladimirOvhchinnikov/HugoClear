package service

import "errors"

// SearchResult представляет собой временную структуру для хранения результатов поиска.
type SearchResult struct {
	// как возможный пример формы ответа
	Result string
}

// SearchService управляет логикой поиска.
type SearchService struct {
}

// SearchServiceOption определяет тип функции опции для SearchController.
type SearchServiceOption func(*SearchService)

// NewSearchService создает новый экземпляр SearchController с применением переданных опций.
func NewSearchService(options ...SearchServiceOption) *SearchService {

	var controller SearchService = SearchService{}

	for _, option := range options {
		option(&controller)
	}

	return &controller
}

// Searcher определяет интерфейс для компонентов, выполняющих поиск.
type Searcher interface {
	Search(query string) (SearchResult, error)
}

// Search выполняет поиск и возвращает результаты.
func (s *SearchService) Search(query string) (SearchResult, error) {

	var search SearchResult = SearchResult{}

	if query != "" {
		search = SearchResult{
			Result: "Все хорошо. Псевдо запрос обработан",
		}
	} else {
		return SearchResult{}, errors.New("query is empty")
	}

	return search, nil
}

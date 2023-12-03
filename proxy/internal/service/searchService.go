package service

import "errors"

type GeoSearch struct {
}

type SearchOption func(*GeoSearch)

type SearchResult struct {
	Result string
}

func NewSearch(options ...SearchOption) *GeoSearch {
	var s GeoSearch
	for _, option := range options {
		option(&s)
	}
	return &s
}

func (s *GeoSearch) Search(query string) (SearchResult, error) {
	if query == "" {
		return SearchResult{}, errors.New("query is empty")
	}
	// Имитация поиска
	return SearchResult{Result: "Все хорошо. Псевдо запрос обработан"}, nil
}

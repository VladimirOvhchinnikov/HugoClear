package service

import "errors"

type Searcher interface {
	Search(query string) (SearchResult, error)
}

type SearchResult struct {
	Result string
}

type Search struct {
}

type SearchOption func(*Search)

func NewSearch(options ...SearchOption) *Search {
	var s Search
	for _, option := range options {
		option(&s)
	}
	return &s
}

func (s *Search) Search(query string) (SearchResult, error) {
	if query == "" {
		return SearchResult{}, errors.New("query is empty")
	}
	// Имитация поиска
	return SearchResult{Result: "Все хорошо. Псевдо запрос обработан"}, nil
}

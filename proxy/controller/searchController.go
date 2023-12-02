package controller

import (
	"log"
	"net/http"
	"proxy/service"
	"proxy/utils"
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
func HandleSearch(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleSearch - work now")

	// Extracting the format of the incoming request
	var extraData utils.ExtractDataFromRequest
	var err error
	extraData, err = extraData.Extract(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Unmarshalling the request body
	var search searchUnm
	err = extraData.UnmarshalAndProcess(r, &search)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Sending data to the business logic layer
	var searchLogic service.SearchService = *service.NewSearchService()
	var searchLogicResult service.SearchResult = service.SearchResult{}
	searchLogicResult, err = searchLogic.Search(search.Query)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Encoding and sending the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(searchLogicResult.Result))
}

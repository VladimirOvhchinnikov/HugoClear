package controller

import (
	"log"
	"net/http"
	"proxy/service"
	"proxy/utils"
)

// geoCodeUnm is a temporary structure for the form. Needs to be refined later.
type geoCodeUnm struct {
	Query string `json:"query"`
}

func (g geoCodeUnm) Process() error {
	return nil
}

// HandleGeoCode handles geocoding requests.
// @Summary Handle geocoding
// @Description This endpoint processes geocoding requests and returns the result.
// @Tags GeoCoding
// @Accept  json
// @Produce  json
// @Param query body geoCodeUnm true "Geocoding Query"
// @Success 200 {string} string "Successfully processed geocoding"
// @Failure 400 {string} string "Bad Request"
// @Router /geocode [post]
func HandleGeoCode(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleGeoCode - work now")

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
	var geocode geoCodeUnm
	err = extraData.UnmarshalAndProcess(r, &geocode)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Sending data to the business logic layer
	var geoLogic service.GeoService = *service.NewGeoService()
	var geoLogicResult service.GeoResult = service.GeoResult{}
	geoLogicResult, err = geoLogic.GeoSE(geocode.Query)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Encoding and sending the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(geoLogicResult.Result))
}

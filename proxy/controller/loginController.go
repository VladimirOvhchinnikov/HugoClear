package controller

import (
	"encoding/json"
	"log"
	"mode/proxy/service"
	"mode/proxy/utils"
	"net/http"
)

// Контроллер слой
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Извлечение и анмаршалинг данных
	var extraData utils.ExtractDataFromRequest
	var err error
	extraData, err = extraData.Extract(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "incorrect data type", http.StatusUnsupportedMediaType)
		return
	}

	var credentials service.Credentials
	err = extraData.UnmarshalAndProcess(r, &credentials)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Делегирование бизнес-логике
	jwtToken, err := service.AuthenticateUser(credentials)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Отправка успешного ответа
	jsonResponse, err := json.Marshal(jwtToken)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

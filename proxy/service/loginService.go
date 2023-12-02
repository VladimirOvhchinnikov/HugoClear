package service

import (
	"encoding/json"
	"errors"
	"mode/proxy/middleware"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c Credentials) Process() error {
	return nil
}

type JwtResponseBody struct {
	Token string `json:"token"`
}

// LogPas определяет структуру входных данных для авторизации пользователя.
// swagger:model
type LogPas struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginController управляет логикой входа в систему.
type LoginController struct {
}

// LoginControllerOption определяет тип функции опции для LoginController.
type LoginControllerOption func(*LoginController)

// NewLoginController создает новый экземпляр LoginController с применением переданных опций.
func NewLoginController(options ...LoginControllerOption) *LoginController {
	var controller LoginController = LoginController{}

	for _, option := range options {
		option(&controller)
	}

	return &controller
}

// Бизнес слой
func AuthenticateUser(credentials Credentials) (string, error) {
	if CheckLogin(credentials.Username) && CheckPassword(credentials.Password) {
		return middleware.JwtCreate(), nil
	}
	return "", errors.New("unauthorized")
}

func SendJwtResponse(w http.ResponseWriter, jwtToken string) {
	response := middleware.JwtResponse{Body: JwtResponseBody{Token: jwtToken}}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func CheckLogin(login string) bool {
	var testLogPas LogPas = LogPas{
		Username: "user123",
	}

	if login == testLogPas.Username {

		return true
	}

	return false
}

func CheckPassword(password string) bool {
	var testLogPas LogPas = LogPas{
		Password: "mypassword",
	}

	if password == testLogPas.Password {

		return true
	}
	return false
}

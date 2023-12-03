package service

import (
	"errors"
	"proxy/middleware"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c Credentials) Process() error {
	return nil
}

// LogPas определяет структуру входных данных для авторизации пользователя.
// swagger:model

// loginController управляет логикой входа в систему.
type loginController struct {
}

// LoginControllerOption определяет тип функции опции для LoginController.
type LoginControllerOption func(*loginController)

// NewLoginController создает новый экземпляр LoginController с применением переданных опций.
func NewLoginController(options ...LoginControllerOption) *loginController {
	var controller loginController = loginController{}

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

func CheckLogin(login string) bool {
	var testLogPas Credentials = Credentials{
		Username: "user123",
	}

	if login == testLogPas.Username {

		return true
	}

	return false
}

func CheckPassword(password string) bool {
	var testLogPas Credentials = Credentials{
		Password: "mypassword",
	}

	if password == testLogPas.Password {

		return true
	}
	return false
}

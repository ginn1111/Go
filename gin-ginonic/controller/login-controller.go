package controller

import (
	"github.com/go/go-goinc/entity"
	"github.com/go/go-goinc/service"
)

type LoginController interface {
	Login(credentials entity.Credentials) string
}

type loginController struct {
	jwtService   service.JWTService
	loginService service.LoginService
}

func NewLoginController(jwtService service.JWTService, loginService service.LoginService) LoginController {
	return &loginController{
		jwtService,
		loginService,
	}
}

func (lc *loginController) Login(credentials entity.Credentials) string {
	lc.loginService = service.NewLoginService()

	isAccountExist := lc.loginService.Login(credentials.Username, credentials.Password)

	if isAccountExist {
		return lc.jwtService.GenerateToken(credentials.Username, credentials.Admin)
	}

	return ""
}

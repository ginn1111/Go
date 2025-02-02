package service

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	username string
	password string
}

func NewLoginService() *loginService {
	return &loginService{
		username: "Gin",
		password: "Thu@n12312",
	}
}

func (ls *loginService) Login(username string, password string) bool {
	return username == ls.username && password == ls.password
}

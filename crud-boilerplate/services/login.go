package services

import (
	loginDomain "github.com/AndrewJoyT/crud-boilerplate/domain/login"
	"github.com/AndrewJoyT/crud-boilerplate/utils/errors"
)

var (
	LoginService loginServiceInterface = &loginService{}
)

type loginService struct {
}

type loginServiceInterface interface {
	Login(*loginDomain.LoginRequest) (interface{}, *errors.RestErr)
}

func (s *loginService) Login(payload *loginDomain.LoginRequest) (interface{}, *errors.RestErr) {
	result, err := payload.Login()
	if err != nil {
		return nil, err
	}
	return result, nil
}

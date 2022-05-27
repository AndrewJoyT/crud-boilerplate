package services

import (
	regisDomain "github.com/AndrewJoyT/crud-boilerplate/domain/register"
	"github.com/AndrewJoyT/crud-boilerplate/utils/errors"
)

var (
	RegisterService registerServiceInterface = &registerService{}
)

type registerService struct {
}

type registerServiceInterface interface {
	CreateUser(*regisDomain.RegisterUserRequest) (interface{}, *errors.RestErr)
	Confirmation(token string) *errors.RestErr
}

func (s *registerService) CreateUser(payload *regisDomain.RegisterUserRequest) (interface{}, *errors.RestErr) {
	result, err := payload.CreateUser()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *registerService) Confirmation(token string) *errors.RestErr {
	if err := regisDomain.Confirmation(token); err != nil {
		return err
	}
	return nil
}

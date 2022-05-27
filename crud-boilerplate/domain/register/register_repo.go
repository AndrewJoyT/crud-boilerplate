package register

import (
	"github.com/AndrewJoyT/crud-boilerplate/database/schema"
	"github.com/AndrewJoyT/crud-boilerplate/utils/errors"

	"github.com/AndrewJoyT/crud-boilerplate/domain/password"
)

func (s *RegisterUserRequest) CreateUserToDB() *errors.RestErr {
	pass := password.HashAndSalt(password.GetPwd(s.Password))
	userData := &schema.Users{
		Name:      s.FirstName + " " + s.LastName,
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Email:     s.Email,
		Password:  pass,
		//StatusID:          1,
		ConfirmationToken: s.ConfirmationToken,
		AuthToken:         "",
		Image:             s.Image,
		Role:              s.Role,
		Nip:               s.Nip,
	}

	if err := schema.Database.Create(userData); err.Error != nil {
		return errors.NewInternalServerError("error when trying to save callback user to db: " + err.Error.Error())
	}
	return nil
}

func GetUserByConfirmationToken(token string) (*schema.Users, *errors.RestErr) {
	var user schema.Users
	if err := schema.Database.First(&user, "confirmation_token = ?", token); err.Error != nil {
		return nil, errors.NewInternalServerError("error when trying to get user by confirmation token: " + err.Error.Error())
	}
	return &user, nil
}

func UpdateUser(user *schema.Users) *errors.RestErr {
	if err := schema.Database.Save(user); err.Error != nil {
		return errors.NewInternalServerError("error when tryin to update users: " + err.Error.Error())
	}
	return nil
}

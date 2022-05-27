package register

type RegisterUserRequest struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmationToken string `json:confirmation_token`
	Role              string `json:"role"`
	Image             string `json:"image"`
	Nip               string `json:"nip"`
}

type RegisterAdminResponse struct {
	Email string `json:"email"`
}

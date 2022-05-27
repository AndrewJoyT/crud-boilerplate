package login

type LoginRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserUUID  string `json:"id"`
	UserNip   string `json:"user_nip"`
	IsSuccess bool   `json:"is_success"`
	Token     string `json:"token"`
}

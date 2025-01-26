package domain

type SignUpRequest struct {
	Firstname string `json:"first_name" binding:"required"`
	Lastname  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type SignUpResponse struct {
	AccessToken  string `json:access_token`
	RefreshToken string `json:refresh_token`
}

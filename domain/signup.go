package domain

type SignUpRequest struct {
	Firstname string `form:"first_name" binding:"required"`
	Lastname  string `form:"last_name" binding:"required"`
	Email     string `form:"email" binding:"required,email"`
	Password  string `form:"password" binding:"required"`
}

type SignUpResponse struct {
	AccessToken  string `json:access_token`
	RefreshToken string `json:refresh_token`
}

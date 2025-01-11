package domain

type SignUpRequest struct {
	Firstname string `json:first_name`
	Lastname  string `json:last_name`
	Email     string `json:email`
	Password  string `json:password`
}

type SignUpResponse struct {
	AccessToken  string `json:access_token`
	RefreshToken string `json:refresh_token`
}



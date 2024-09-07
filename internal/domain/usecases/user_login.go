package domain

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type UserLoginResponse struct {
	AccessToken string `json:"accessToken"`
}

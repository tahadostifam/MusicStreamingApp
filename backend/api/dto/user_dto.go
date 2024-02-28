package dto

type UserSigninDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"Password" validate:"required"`
}

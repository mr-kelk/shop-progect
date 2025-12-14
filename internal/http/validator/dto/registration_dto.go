package dto

type RegistrationRequest struct {
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
	Login    string `json:"login" validate:"required,min=3,max=32"`
}

package dto

type RegisterDTO struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=200"`
	LastName  string `json:"last_name" validate:"required,min=2,max=200"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=128"`
}

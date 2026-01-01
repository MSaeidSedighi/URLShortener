package dto

type UpdateUserDTO struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=200"`
	LastName  string `json:"last_name" validate:"required,min=2,max=200"`
}

package dto

type PlayerDTO struct {
	ID             int    `json:"id"`
	FirstName      string `json:"first_name" validate:"required"`
	LastName       string `json:"last_name,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
}

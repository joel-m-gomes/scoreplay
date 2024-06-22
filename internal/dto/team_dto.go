package dto

type TeamDTO struct {
	ID      int         `json:"id"`
	Name    string      `json:"name" validate:"required"`
	Logo    string      `json:"logo,omitempty"`
	Players []PlayerDTO `json:"players,omitempty"`
}

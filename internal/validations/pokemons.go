package validations

import (
	"time"

	"github.com/google/uuid"
)

type CreatePokemonRequest struct {
	Name   string `validate:"required, min=2,max=30" json:"name"`
	Type   string `validate:"required, min=2,max=10" json:"type"`
	Region string `validate:"required, min=2,max=20" json:"region"`
	Image  string `validate:"required, min=10,max=255" json:"image"`
}

type UpdatePokemonRequest struct {
	Id        uuid.UUID `validate:"required"`
	Name      string    `validate:"required, min=2,max=30" json:"name"`
	Type      string    `validate:"required, min=2,max=10" json:"type"`
	Region    string    `validate:"required, min=2,max=20" json:"region"`
	Image     string    `validate:"required, min=10,max=255" json:"image"`
	UpdatedAt time.Time
}

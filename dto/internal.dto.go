package dto

import (
	"github.com/google/uuid"
)

type BlackjackInterimMessageDTO struct {
	UserId  uuid.UUID `json:"userId"`
	Enabled bool      `json:"enabled"`
}

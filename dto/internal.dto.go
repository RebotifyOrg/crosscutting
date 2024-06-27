package dto

import (
	"encoding/json"

	"github.com/google/uuid"
)

type BlackjackInterimMessageDTO struct {
	UserId  uuid.UUID `json:"userId"`
	Enabled bool      `json:"enabled"`
}

func (i BlackjackInterimMessageDTO) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i *BlackjackInterimMessageDTO) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, i)
}
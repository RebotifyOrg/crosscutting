package dto

import (
	"github.com/RebotifyOrg/crosscutting/internalevent"
	"github.com/google/uuid"
)

type BlackjackInterimMessageDTO struct {
	Event internalevent.InternalEvent    `json:"event"`
	Data  BlackjackInterimMessageDataDTO `json:"data"`
}
type BlackjackInterimMessageDataDTO struct {
	UserId  uuid.UUID `json:"userId"`
	Enabled bool      `json:"enabled"`
}

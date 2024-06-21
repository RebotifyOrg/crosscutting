package dto

import (
	"github.com/RebotifyOrg/crosscutting/gametype"
	"github.com/google/uuid"
)

type GameServerMessageDTO struct {
	GameType gametype.GameType `json:"gameType"`
	UserId   uuid.UUID         `json:"userId"`
	Data     interface{}       `json:"data"`
}

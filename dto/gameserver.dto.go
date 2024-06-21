package dto

import (
	"github.com/RebotifyOrg/crosscutting/gametype"
	"github.com/RebotifyOrg/crosscutting/gametypeevent"
	"github.com/google/uuid"
)

type GameServerMessageDTO struct {
	GameType gametype.GameType        `json:"gameType"`
	UserId   uuid.UUID                `json:"userId"`
	Data     GameServerMessageDataDTO `json:"data"`
}
type GameServerMessageDataDTO struct {
	Event  gametypeevent.GameTypeEvent `json:"event"`
	UserId uuid.UUID                   `json:"userId"`
}

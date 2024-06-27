package dto

import (
	"encoding/json"

	"github.com/RebotifyOrg/crosscutting/gametype"
	"github.com/RebotifyOrg/crosscutting/gametypeevent"
	"github.com/google/uuid"
)

type GameServerMessageDTO struct {
	GameType gametype.GameType        `json:"gameType"`
	UserId   uuid.UUID                `json:"userId"`
	Data     GameServerMessageDataDTO `json:"data"`
}

func (i GameServerMessageDTO) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i *GameServerMessageDTO) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, i)
}

type GameServerMessageDataDTO struct {
	Event    gametypeevent.GameTypeEvent `json:"event"`
	UserId   uuid.UUID                   `json:"userId"`
	Username string                      `json:"username"`
}

func (i GameServerMessageDataDTO) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i *GameServerMessageDataDTO) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, i)
}

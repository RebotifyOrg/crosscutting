package dto

import (
	"encoding/json"

	"github.com/RebotifyOrg/crosscutting/internalevent"
	"github.com/RebotifyOrg/crosscutting/platform"
)

type EventProcessorMessageDTO struct {
	Platform platform.Platform `json:"platform"`
	Data     json.RawMessage   `json:"data"`
}

type EventProcessorMessageDataDTO struct {
	Event internalevent.InternalEvent `json:"event"`
	Data  interface{}                 `json:"data"`
}

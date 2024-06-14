package dto

import (
	"github.com/RebotifyOrg/crosscutting/internalevent"
	"github.com/RebotifyOrg/crosscutting/platform"
)

type EventProcessorMessageDTO struct {
	Platform platform.Platform `json:"platform"`
	Data     interface{}       `json:"data"`
}

type EventProcessorMessageDataDTO struct {
	Event internalevent.InternalEvent `json:"event"`
	Data  interface{}                 `json:"data"`
}

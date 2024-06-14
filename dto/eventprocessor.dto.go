package dto

import "github.com/RebotifyOrg/crosscutting/platform"

type EventProcessorMessageDTO struct {
	Platform platform.Platform `json:"platform"`
	Data     interface{}       `json:"data"`
}

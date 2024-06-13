package dto

import "github.com/ReBotifyOrg/crosscutting/platform"

type EventProcessorMessageDTO struct {
	Platform platform.Platform `json:"platform"`
	Data     []byte            `json:"data"`
}

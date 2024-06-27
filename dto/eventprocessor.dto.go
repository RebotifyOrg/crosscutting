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
	Data  json.RawMessage             `json:"data"`
}

func (i EventProcessorMessageDTO) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i *EventProcessorMessageDTO) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, i)
}

func (i EventProcessorMessageDataDTO) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i *EventProcessorMessageDataDTO) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, i)
}
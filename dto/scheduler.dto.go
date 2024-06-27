package dto

import (
	"encoding/json"

	"github.com/RebotifyOrg/crosscutting/internalevent"
	"github.com/google/uuid"
)

type SchedulerDTO struct {
	ScheduleId    uint64                      `json:"scheduleId"`
	InternalEvent internalevent.InternalEvent `json:"internalEvent"`
	ScheduleEvent int                         `json:"scheduleEvent"`
	UserId        uuid.UUID                   `json:"userId"`
	Seconds       int                         `json:"seconds"`
}

func (i SchedulerDTO) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}
func (i *SchedulerDTO) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, i)
}

package dto

import (
	"encoding/json"

	"github.com/RebotifyOrg/crosscutting/internalevent"
	"github.com/google/uuid"
)

type SchedulerDTO struct {
	InternalEvent internalevent.InternalEvent `redis:"internalEvent"`
	ScheduleEvent int                         `redis:"scheduleEvent"`
	UserId        uuid.UUID                   `redis:"userId"`
	Seconds       int                         `redis:"seconds"`
}

func (i SchedulerDTO) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}
func (i *SchedulerDTO) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, i)
}

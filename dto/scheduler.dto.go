package dto

import (
	"github.com/RebotifyOrg/crosscutting/internalevent"
	"github.com/google/uuid"
)

type SchedulerDTO struct {
	InternalEvent internalevent.InternalEvent `redis:"internalEvent"`
	ScheduleEvent int                         `redis:"scheduleEvent"`
	UserId        uuid.UUID                   `redis:"userId"`
	Seconds       int                         `redis:"seconds"`
}

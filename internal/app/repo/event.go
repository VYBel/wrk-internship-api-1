package repo

import (
	"github.com/ozonmp/wrk-internship-api/internal/model"
)

type EventRepo interface {
	Lock(n uint64) ([]model.InternshipEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.InternshipEvent) error
	Remove(eventIDs []uint64) error
}

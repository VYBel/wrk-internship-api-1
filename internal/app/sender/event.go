package sender

import (
	"github.com/ozonmp/wrk-internship-api/internal/model"
)

type EventSender interface {
	Send(internship *model.InternshipEvent) error
}

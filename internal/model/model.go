package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID         primitive.ObjectID `bson:"_id"`
	Type       string             `bson:"type"`        // тип события
	State      int                `bson:"state"`       // state 0 - для незавершенных событий, 1 - для завершенных
	StartedAt  time.Time          `bson:"started_at"`  // время начала события
	FinishedAt time.Time          `bson:"finished_at"` // время завершения события
}

type Request struct {
	Type string `json:"type"`
}

func (r *Request) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.Type, validation.Required, validation.By(required())),
	)
}

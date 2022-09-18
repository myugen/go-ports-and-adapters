package uuid

import (
	"github.com/google/uuid"
)

type UUID struct {
	uuid.UUID
}

func MustParse(value string) UUID {
	return UUID{uuid.MustParse(value)}
}

var Nil = UUID{uuid.Nil}

func New() UUID {
	return UUID{uuid.New()}
}

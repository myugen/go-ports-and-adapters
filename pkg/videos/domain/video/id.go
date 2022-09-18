package video

import "github.com/myugen/go-ports-and-adapters/internal/shared/domain/valueobjects/uuid"

type ID = uuid.UUID

func NewID(u ...uuid.UUID) ID {
	if len(u) == 1 {
		return ID{UUID: u[0].UUID}
	}

	return ID{UUID: uuid.New().UUID}
}

var NoID = NewID(uuid.Nil)

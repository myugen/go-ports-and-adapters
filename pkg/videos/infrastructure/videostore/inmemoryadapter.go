package videostore

import (
	"fmt"

	"github.com/myugen/go-ports-and-adapters/pkg/videos/domain/video"
)

var DuplicatedTitleErrorMsgTemplate = "a video with same title %q exists"

type InMemoryStore map[video.ID]video.Video

type InMemoryAdapter struct {
	store InMemoryStore
}

func (m *InMemoryAdapter) Create(videoToCreate video.Video) error {
	for _, value := range m.store {
		if value.Title == videoToCreate.Title {
			return fmt.Errorf(DuplicatedTitleErrorMsgTemplate, videoToCreate.Title)
		}
	}
	m.store[videoToCreate.ID] = videoToCreate
	return nil
}

func (m *InMemoryAdapter) NextID() video.ID {
	return video.NewID()
}

func NewInMemoryAdapter(store ...InMemoryStore) *InMemoryAdapter {
	s := InMemoryStore{}
	if len(store) == 1 {
		s = store[0]
	}
	return &InMemoryAdapter{store: s}
}

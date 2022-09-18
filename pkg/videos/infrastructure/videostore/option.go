package videostore

import "github.com/myugen/go-ports-and-adapters/pkg/videos/application/usecases/videocreator"

func WithInMemoryRepository(store ...InMemoryStore) videocreator.Option {
	repository := NewInMemoryAdapter(store...)
	return videocreator.WithRepository(repository)
}

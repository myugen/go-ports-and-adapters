package videocreator

import "github.com/myugen/go-ports-and-adapters/pkg/videos/domain/video"

type Option func(*VideoCreator) error

func WithRepository(repository video.Repository) Option {
	return func(videoCreator *VideoCreator) error {
		videoCreator.repository = repository
		return nil
	}
}

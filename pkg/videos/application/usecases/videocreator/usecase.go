package videocreator

import (
	"github.com/myugen/go-ports-and-adapters/pkg/videos/domain/video"
)

type Command struct {
	Title video.Title
}

type VideoCreator struct {
	repository video.Repository
}

func (u VideoCreator) Invoke(command Command) (video.ID, error) {
	nextID := u.repository.NextID()
	newVideo := video.Video{
		ID:    nextID,
		Title: command.Title,
	}
	if err := u.repository.Create(newVideo); err != nil {
		return video.NoID, err
	}
	return newVideo.ID, nil
}

func New(options ...Option) *VideoCreator {
	videoCreator := &VideoCreator{}
	for _, option := range options {
		if err := option(videoCreator); err != nil {
			panic(err)
		}
	}
	return videoCreator
}

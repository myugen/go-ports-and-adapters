package video

//go:generate mockgen -destination=../../../../mocks/pkg/videos/domain/video/repository.go -package=mockvideo . Repository

type Repository interface {
	Create(video Video) error
	NextID() ID
}

package video

type Title = string

type Video struct {
	ID    ID
	Title Title
}

var NoVideo = Video{}

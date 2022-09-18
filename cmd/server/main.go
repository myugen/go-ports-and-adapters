package main

import (
	"fmt"

	"github.com/myugen/go-ports-and-adapters/internal/shared/domain/valueobjects/uuid"
	"github.com/myugen/go-ports-and-adapters/pkg/videos/domain/video"
)

func main() {
	fmt.Println("Initializing server")

	video := video.Video{
		ID:    uuid.New(),
		Title: "foo",
	}

	fmt.Printf("Video created: %v", video)
}

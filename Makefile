.PHONY: mocks

mocks: pkg/videos/domain/video/repository.go
	@echo "Generating mocks..."
	@go generate ./...

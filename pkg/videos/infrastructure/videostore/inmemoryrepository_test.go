package videostore_test

import (
	"fmt"
	"testing"

	"github.com/myugen/go-ports-and-adapters/pkg/videos/domain/video"
	"github.com/myugen/go-ports-and-adapters/pkg/videos/infrastructure/videostore"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryRepository_Create_should(t *testing.T) {
	type fields struct {
		store videostore.InMemoryStore
	}
	type args struct {
		video video.Video
	}
	type want struct {
		error error
	}
	testcases := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name:   "return no error when video is correct",
			fields: fields{store: videostore.InMemoryStore{}},
			args: args{video: video.Video{
				ID:    video.NewID(),
				Title: "Test",
			}},
			want: want{error: nil},
		},
		{
			name: "return an error when video title is duplicated",
			fields: fields{store: videostore.InMemoryStore{video.NewID(): {
				ID:    video.NoID,
				Title: "Duplicated Title",
			}}},
			args: args{video: video.Video{
				ID:    video.NewID(),
				Title: "Duplicated Title",
			}},
			want: want{error: fmt.Errorf(videostore.DuplicatedTitleErrorMsgTemplate, "Duplicated Title")},
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			repositoryAdapter := videostore.NewInMemoryAdapter(testcase.fields.store)
			err := repositoryAdapter.Create(testcase.args.video)
			assert.Equal(t, testcase.want.error, err)
		})
	}
}

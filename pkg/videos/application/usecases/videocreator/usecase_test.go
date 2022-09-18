package videocreator_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/myugen/go-ports-and-adapters/internal/shared/domain/valueobjects/uuid"
	mockvideo "github.com/myugen/go-ports-and-adapters/mocks/pkg/videos/domain/video"
	"github.com/myugen/go-ports-and-adapters/pkg/videos/application/usecases/videocreator"
	"github.com/myugen/go-ports-and-adapters/pkg/videos/domain/video"
	"github.com/stretchr/testify/assert"
)

func TestVideoCreator_Invoke_should(t *testing.T) {
	mockUUID := uuid.MustParse("0956a14e-fb41-4d5b-8ac1-53aefbe5bcc3")
	type fields struct {
		repository *mockvideo.MockRepository
	}
	type args struct {
		command videocreator.Command
	}
	type want struct {
		id    video.ID
		error error
	}
	testcases := []struct {
		name    string
		prepare func(f *fields)
		args    args
		want    want
	}{
		{
			name: "return the ID of created video and no error",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.repository.EXPECT().NextID().Return(mockUUID),
					f.repository.EXPECT().Create(video.Video{ID: mockUUID, Title: "Test Title"}).Return(nil),
				)
			},
			args: args{command: videocreator.Command{Title: "Test Title"}},
			want: want{
				id:    mockUUID,
				error: nil,
			},
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			f := &fields{
				repository: mockvideo.NewMockRepository(mockCtrl),
			}
			if testcase.prepare != nil {
				testcase.prepare(f)
			}
			usecase := videocreator.New(videocreator.WithRepository(f.repository))
			got, err := usecase.Invoke(testcase.args.command)
			assert.Equal(t, testcase.want.id, got)
			assert.Equal(t, testcase.want.error, err)
		})
	}
}

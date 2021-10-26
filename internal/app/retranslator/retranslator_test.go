package retranslator

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/wrk-internship-api/internal/mocks"
)

func TestStart(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	repo.EXPECT().Lock(gomock.Any()).AnyTimes()

	cfg := Config{
		ChannelSize:    512,
		ConsumerCount:  2,
		ConsumeSize:    10,
		ConsumeTimeout: time.Second,
		ProducerCount:  2,
		WorkerCount:    2,
		Repo:           repo,
		Sender:         sender,
	}

	r := NewRetranslator(cfg)
	r.Start()
	time.Sleep(2 * time.Second)
	r.Close()
}

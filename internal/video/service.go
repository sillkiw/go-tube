package video

import (
	"gotube/internal/config"
	"log/slog"
	"sync/atomic"
)

type Service struct {
	logger      *slog.Logger
	cfg         *config.Config
	taskCh      chan task
	uploadCount int64
	queueLen    int32
}

func NewService(logger *slog.Logger, cfg *config.Config) *Service {
	s := &Service{
		logger:      logger,
		cfg:         cfg,
		taskCh:      make(chan task, 50),
		uploadCount: 0,
		queueLen:    0,
	}
	go s.convertVideo()
	return s
}
func (s *Service) IncrementUploaded() {
	atomic.AddInt64(&s.uploadCount, 1)
}

func (s *Service) GetUploadedCount() int64 {
	return atomic.LoadInt64(&s.uploadCount)
}

func (s *Service) addToQueue(n int) {
	atomic.AddInt32(&s.queueLen, int32(n))
}

func (s *Service) DecreaseQueue(n int) {
	atomic.AddInt32(&s.queueLen, -int32(n))
}

func (s *Service) GetQueueLen() int32 {
	return atomic.LoadInt32(&s.queueLen)
}

package service

import (
	"context"

	"github.com/aikizoku/rabbitgo/appengine/src/lib/log"
	"github.com/aikizoku/rabbitgo/appengine/src/repository"
)

type sample struct {
	repo repository.Sample
}

func (s *sample) Sample(ctx context.Context) error {
	err := s.repo.Sample(ctx)
	if err != nil {
		log.Errorm(ctx, "s.repo.Sample", err)
		return err
	}
	return nil
}

// NewSample ... サービスを作成する
func NewSample(repo repository.Sample) Sample {
	return &sample{
		repo: repo,
	}
}

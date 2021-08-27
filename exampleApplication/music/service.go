package music

import (
	"context"
)

type SongStorage interface {
	GetTrack(context.Context, int) (Track, error)
	SaveTrack(context.Context, Track) error
}

type Service struct {
	storage SongStorage
}

func NewService(storage SongStorage) *Service {
	s := &Service{
		storage: storage,
	}
	return s
}

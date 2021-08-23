package music

import (
	"context"
	randrand "math/rand"
	"time"
)

type Track struct {
	Id       int           `json:"id"`
	Title    string        `json:"title"`
	Duration time.Duration `json:"duration"`
	Artist   string        `json:"artist"`
}

func (t Track) isValid() bool {
	return true
}

func (s *Service) GetTrack(ctx context.Context, id int) (Track, error) {
	return s.storage.GetTrack(ctx, id)
}

func (s *Service) AddTrack(ctx context.Context, track Track) (Track, error) {
	track.Id = randrand.Int()
	if !track.isValid() {
		return Track{}, ErrorTrackInvalid{}
	}
	// Eventuelle service-greier
	return track, s.storage.SaveTrack(ctx, track)
}

type ErrorTrackInvalid struct{}

func (e ErrorTrackInvalid) Error() string {
	return "Invalid track"
}

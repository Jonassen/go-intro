package storage

import (
	"context"
	"music/music"
)

func NewMemStore() *Memory {
	return &Memory{
		tracks: make(map[int]*music.Track),
	}
}

type Memory struct{
	tracks map[int]*music.Track
}

func (m *Memory) GetTrack(ctx context.Context, i int) (music.Track, error) {
	track := m.tracks[i]
	if track == nil {
		return music.Track{}, ErrorNotFound{}
	}
	return *track, nil
}

func (m *Memory) SaveTrack(_ context.Context, track music.Track) error {
	_, ok := m.tracks[track.Id]
	if ok {
		return ErrorIdCollision{}
	}
	m.tracks[track.Id] = &track
	return nil
}

// Denne kan ikke brukes i pakken songs!
type ErrorNotFound struct{}

func (e ErrorNotFound) Error() string {
	return "Track not found"
}

type ErrorIdCollision struct{}

func (e ErrorIdCollision) Error() string {
	return "Id already in use"
}



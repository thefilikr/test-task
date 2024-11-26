package service

import (
	"fmt"
	"log/slog"
	"test-task-filikr/internal/store"
)

type SongService struct {
	SongStore store.SongStore
	Log       *slog.Logger
}

func (s *SongService) CreateSong(group, song string, releaseDate *string, text *[]string, link *string) error {
	return s.SongStore.CreateSong(group, song, releaseDate, text, link)
}

func (s *SongService) GetSong(group, song, countVerses, numPage string) (*store.Song, error) {
	return s.SongStore.GetSong(group, song, countVerses, numPage)
}

func (s *SongService) GetIDSong(group, song string) (*uint, error) {
	return s.SongStore.GetIDSong(group, song)
}

func (s *SongService) DeleteSong(group, song string) error {
	id, err := s.GetIDSong(group, song)
	if err != nil {
		return fmt.Errorf("the song was not found")
	}

	s.Log.Debug("Successful get id song: ", *id)

	return s.SongStore.DeleteSong(*id)
}

func (s *SongService) EditSong(group, song string, releaseDate *string, text *[]string, link *string) error {
	id, err := s.GetIDSong(group, song)
	if err != nil {
		return fmt.Errorf("the song was not found")
	}

	s.Log.Debug("Successful get id song: ", *id)

	return s.SongStore.EditSong(*id, group, song, releaseDate, text, link)
}
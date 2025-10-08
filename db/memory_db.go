package db

import (
	"movies/entity/character"
	"movies/entity/movie"

	"github.com/google/uuid"
)

type MemoryDb struct {
	Movies map[uuid.UUID]movie.Movie
	Characters map[uuid.UUID]character.Character
}

func New() *MemoryDb {
	return &MemoryDb{
		Movies: make(map[uuid.UUID]movie.Movie),
		Characters: make(map[uuid.UUID]character.Character),
	}
}
package character_repository

import (
	"movies/internal/db"
	"movies/internal/entity/movie"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewMovie(t *testing.T) {
	db := db.New()
	cr := New(db)
	movie := GetMovie(uuid.New(), 2000, "test title")
	name := "test name"

	testCharacter := cr.Create(name, movie)

	assert.Equal(t, 1, getMapLength(&db.Characters))
	assert.Equal(t, name, testCharacter.Name)
	assert.Equal(t, movie, testCharacter.Movie)
}

func TestGetMovie(t *testing.T) {
	db := db.New()
	cr := New(db)
	movie := GetMovie(uuid.New(), 2000, "test title")
	name := "test name"

	testCharacter := cr.Create(name, movie)

	characterGet := cr.Get(testCharacter.Id)
	characterGetEmpty := cr.Get(uuid.New())

	assert.NotEmpty(t, characterGet)
	assert.Equal(t, name, characterGet.Name)
	assert.Equal(t, movie, characterGet.Movie)

	assert.Empty(t, characterGetEmpty)
}

func TestUpdateMovie(t *testing.T) {
	db := db.New()
	cr := New(db)
	movie := GetMovie(uuid.New(), 2000, "test title")
	name := "test name"

	testCharacter := cr.Create(name, movie)

	updatedMovie := GetMovie(uuid.New(), 2025, "updated title")
	updatedCharacter, error := cr.Update(testCharacter.Id, "updated", updatedMovie)

	assert.Nil(t, error)
	assert.Equal(t, "updated", updatedCharacter.Name)
	assert.Equal(t, updatedMovie, updatedCharacter.Movie)
}

func TestDeleteMovie(t *testing.T) {
	db := db.New()
	cr := New(db)
	movie := GetMovie(uuid.New(), 2000, "test title")
	name := "test name"

	c := cr.Create(name, movie)

	cr.Delete(c.Id)

	assert.Equal(t, 0, getMapLength(&db.Characters))
}

func GetMovie(id uuid.UUID, year int, title string) movie.Movie {
	return movie.Movie{Id: id, Year: year, Title: title}
}

func getMapLength(m *sync.Map) int {
	length := 0
	m.Range(func(key, value any) bool {
		length++
		return true 
	})

	return length
}
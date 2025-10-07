package character

import (
	"movies/movie"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMovie(t *testing.T) {
	movie := GetMovie(1, 2000, "test title")
	name := "test name"

	testCharacter := AddCharacter(movie, name)

	assert.Equal(t, 1, len(Characters))
	assert.Equal(t, name, testCharacter.Name)
	assert.Equal(t, movie, testCharacter.Movie)

	ClearCharacters()
}

func TestGetMovie(t *testing.T) {
	movie := GetMovie(1, 2000, "test title")
	name := "test name"

	testCharacter := AddCharacter(movie, name)

	characterGet := Get(testCharacter.Id)
	characterGetEmpty := Get(testCharacter.Id + 1)

	assert.NotEmpty(t, characterGet)
	assert.Equal(t, name, characterGet.Name)
	assert.Equal(t, movie, characterGet.Movie)

	assert.Empty(t, characterGetEmpty)

	ClearCharacters()
}

func TestUpdateMovie(t *testing.T) {
	movie := GetMovie(1, 2000, "test title")
	name := "test name"

	testCharacter := AddCharacter(movie, name)

	updatedMovie := GetMovie(2, 2025, "updated title")
	updatedCharacter := testCharacter.Update(Config{Name: "updated", Movie: updatedMovie})

	assert.Equal(t, "updated", updatedCharacter.Name)
	assert.Equal(t, updatedMovie, updatedCharacter.Movie)

	assert.Equal(t, "updated", testCharacter.Name)
	assert.Equal(t, updatedMovie, testCharacter.Movie)

	ClearCharacters()
}

func TestDeleteMovie(t *testing.T) {
	movie := GetMovie(1, 2000, "test title")
	name := "test name"

	AddCharacter(movie, name)

	Delete(Characters[0].Id)

	assert.Equal(t, 0, len(Characters))

	ClearCharacters()
}

func GetMovie(id int32, year int, title string) movie.Movie {
	return movie.Movie{Id: id, Year: year, Title: title}
}

func AddCharacter(movie movie.Movie, name string) Character {
	cfg := Config{Name: name, Movie: movie}
	return New(cfg)
}

func ClearCharacters() {
	Characters = slices.Delete(Characters, 0, len(Characters))
}
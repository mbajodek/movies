package movie

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMovie(t *testing.T) {
	title := "test title"
	year := 2025

	testMovie := AddMovie(year, title)

	assert.Equal(t, 1, len(Movies))
	assert.Equal(t, title, testMovie.Title)
	assert.Equal(t, year, testMovie.Year)

	ClearMovies()
}

func TestGetMovie(t *testing.T) {
	title := "test title"
	year := 2025

	testMovie := AddMovie(year, title)

	movieGet := Get(testMovie.Id)
	movieGetEmpty := Get(testMovie.Id + 1)

	assert.NotEmpty(t, movieGet)
	assert.Equal(t, title, movieGet.Title)
	assert.Equal(t, year, movieGet.Year)

	assert.Empty(t, movieGetEmpty)

	ClearMovies()
}

func TestUpdateMovie(t *testing.T) {
	title := "test title"
	year := 2025

	testMovie := AddMovie(year, title)

	updatedMovie := testMovie.Update(Config{Title: "updated", Year: 2000})

	assert.Equal(t, "updated", updatedMovie.Title)
	assert.Equal(t, 2000, updatedMovie.Year)

	assert.Equal(t, "updated", testMovie.Title)
	assert.Equal(t, 2000, testMovie.Year)

	ClearMovies()
}

func TestDeleteMovie(t *testing.T) {
	AddMovie(2000, "test 1")
	AddMovie(2001, "test 2")
	AddMovie(2002, "test 3")

	Delete(Movies[2].Id)
	Delete(Movies[1].Id)

	assert.Equal(t, 1, len(Movies))

	ClearMovies()
}

func AddMovie(year int, title string) Movie {
	cfg := Config{Year: year, Title: title}
	return New(cfg)
}

func ClearMovies() {
	Movies = slices.Delete(Movies, 0, len(Movies))
}

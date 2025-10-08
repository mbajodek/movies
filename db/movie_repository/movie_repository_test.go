package movie_repository

import (
	"movies/db"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewMovie(t *testing.T) {
	db := db.New()
	mr := New(db)
	title := "test title"
	year := 2025

	testMovie := mr.Create(title, year)

	assert.Equal(t, 1, len(db.Movies))
	assert.Equal(t, title, testMovie.Title)
	assert.Equal(t, year, testMovie.Year)
}

func TestGetMovie(t *testing.T) {
	db := db.New()
	mr := New(db)
	title := "test title"
	year := 2025

	testMovie := mr.Create(title, year)

	movieGet := mr.Get(testMovie.Id)
	movieGetEmpty := mr.Get(uuid.New())

	assert.NotEmpty(t, movieGet)
	assert.Equal(t, title, movieGet.Title)
	assert.Equal(t, year, movieGet.Year)

	assert.Empty(t, movieGetEmpty)
}

func TestUpdateMovie(t *testing.T) {
	db := db.New()
	mr := New(db)
	title := "test title"
	year := 2025

	testMovie := mr.Create(title, year)

	updatedMovie, error := mr.Update(testMovie.Id, "updated", 2000)

	assert.Nil(t, error)
	assert.Equal(t, "updated", updatedMovie.Title)
	assert.Equal(t, 2000, updatedMovie.Year)
}

func TestDeleteMovie(t *testing.T) {
	db := db.New()
	mr := New(db)
	mr.Create("test 1", 2000)
	m2 := mr.Create("test 2", 2001)
	m3 := mr.Create("test 3", 2002)

	mr.Delete(m2.Id)
	mr.Delete(m3.Id)

	assert.Equal(t, 1, len(db.Movies))
}

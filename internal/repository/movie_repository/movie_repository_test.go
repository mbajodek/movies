package movie_repository

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"movies/internal/db"
	"movies/internal/entity/movie"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewMovie(t *testing.T) {
	db := db.New()
	mr := New(db)
	title := "test title"
	year := 2025
	movie := getTestMovie(title, year)

	testMovie := mr.Create(movie)

	assert.Equal(t, 1, getMapLength(&db.Movies))
	assert.Equal(t, title, testMovie.Title)
	assert.Equal(t, year, testMovie.Year)
}

func TestGetMovie(t *testing.T) {
	db := db.New()
	mr := New(db)
	title := "test title"
	year := 2025
	movie := getTestMovie(title, year)

	testMovie := mr.Create(movie)

	movieGet, _ := mr.Get(testMovie.Id)
	movieGetEmpty, _ := mr.Get(uuid.New())

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
	movie := getTestMovie(title, year)

	testMovie := mr.Create(movie)

	updatedMovie, error := mr.Update(testMovie.Id, "updated", 2000)

	assert.Nil(t, error)
	assert.Equal(t, "updated", updatedMovie.Title)
	assert.Equal(t, 2000, updatedMovie.Year)
}

func TestDeleteMovie(t *testing.T) {
	db := db.New()
	mr := New(db)
	mr.Create(getTestMovie("test 1", 2000))
	m2 := mr.Create(getTestMovie("test 2", 2001))
	m3 := mr.Create(getTestMovie("test 3", 2002))

	mr.Delete(m2.Id)
	mr.Delete(m3.Id)

	assert.Equal(t, 1, getMapLength(&db.Movies))
}

func getTestMovie(title string, year int) movie.Movie {
	cert := x509.Certificate{}
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	return *movie.NewWithOptions(
		movie.WithTitle(title),
		movie.WithYear(year),
		movie.WithCert(&cert, priv),
	)
}

func getMapLength(m *sync.Map) int {
	length := 0
	m.Range(func(key, value any) bool {
		length++
		return true 
	})

	return length
}

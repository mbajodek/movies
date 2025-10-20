package movie_repository

import (
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"
	"movies/internal/db"
	"movies/internal/entity/movie"

	"github.com/google/uuid"
)

type MovieRepository struct {
	DB *db.MemoryDb
}

func New(db *db.MemoryDb) *MovieRepository {
	return &MovieRepository{DB: db}
}

func (mr *MovieRepository) Create(title string, year int, cert *x509.Certificate, privateKey *rsa.PrivateKey) movie.Movie {
	fmt.Println(title, year)
	movie := movie.NewWithOptions(
		movie.WithTitle(title),
		movie.WithYear(year),
		movie.WithCert(cert, privateKey),
	)

	fmt.Println(movie)
	mr.DB.Movies.Store(movie.Id, *movie)
	return *movie
}

func (mr *MovieRepository) Get(id uuid.UUID) (movie.Movie, bool) {
	v, ok :=  mr.DB.Movies.Load(id)
	var m movie.Movie

	if !ok {
		fmt.Println("No movie with id:", id)
		return m, false
	}
	m = v.(movie.Movie)

	return m, true
}

func (mr *MovieRepository) GetAll() []movie.Movie {
	var movies []movie.Movie
	mr.DB.Movies.Range(func(key, value interface{}) bool {
        movies = append(movies, value.(movie.Movie))
        return true
    })

	return movies
}

func (mr *MovieRepository) Update(id uuid.UUID, title string, year int) (movie.Movie, error) {
	v, ok :=  mr.DB.Movies.Load(id)
	var m movie.Movie
	
	if !ok {
		fmt.Println("No movie with id:", id)
		return m, errors.New("No movie with id: " + id.String())
	}
	m = v.(movie.Movie)
	m.Year = year
	m.Title = title
	mr.DB.Movies.Swap(id, m)

	return m, nil
}

func (mr *MovieRepository) Delete(id uuid.UUID) {
	_, ok :=  mr.DB.Movies.Load(id)
	
	if !ok {
		fmt.Println("No movie with id:", id)
	} else {
		mr.DB.Movies.Delete(id)
	}	
}
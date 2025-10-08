package movie_repository

import (
	"errors"
	"fmt"
	"movies/db"
	"movies/entity/movie"

	"github.com/google/uuid"
)

type MovieRepository struct {
	DB *db.MemoryDb
}

func New(db *db.MemoryDb) *MovieRepository {
	return &MovieRepository{DB: db}
}

func (mr *MovieRepository) Create(title string, year int) movie.Movie {
	movie := movie.NewWithOptions(
		movie.WithTitle(title),
		movie.WithYear(year),
	)

	mr.DB.Movies[movie.Id] = *movie
	return *movie
}

func (mr *MovieRepository) Get(id uuid.UUID) movie.Movie {
	movie, exists :=  mr.DB.Movies[id]
	
	if !exists {
		fmt.Println("No movie with id:", id)
		fmt.Println(movie)
		return movie
	}

	return movie
}

func (mr *MovieRepository) Update(id uuid.UUID, title string, year int) (movie.Movie, error) {
	movie, exists :=  mr.DB.Movies[id]
	
	if !exists {
		fmt.Println("No movie with id:", id)
		return movie, errors.New("No movie with id: " + id.String())
	}
	movie.Year = year
	movie.Title = title
	mr.DB.Movies[id] = movie

	return movie, nil
}

func (mr *MovieRepository) Delete(id uuid.UUID) {
	_, exists :=  mr.DB.Movies[id]
	
	if !exists {
		fmt.Println("No movie with id:", id)
	}

	delete(mr.DB.Movies, id)
}
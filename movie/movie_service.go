package movie

import (
	"fmt"
	"math/rand"
	"slices"
)

func New(cfg Config) Movie {
	newMovie := Movie{
		Id: rand.Int31(),
		Year: cfg.Year,
		Title: cfg.Title,
	}
	Movies = append(Movies, newMovie)

	return newMovie
}

func Get(id int32) Movie {
	for i, movie := range Movies {
		if movie.Id == id {
			return Movies[i]
		}
	}
	fmt.Println("No movie with id:", id)

	return Movie{}
}

func (movie *Movie) Update(cfg Config) Movie {
	movie.Year = cfg.Year
	(*movie).Title = cfg.Title

	return *movie
}

func Delete(id int32) {
	for i, movie := range Movies {
		if movie.Id == id {
			Movies = slices.Delete(Movies, i, i+1)
			break
		}
	}
}
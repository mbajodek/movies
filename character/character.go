package character

import (
	"movies/movie"
)

type Character struct {
	Id int32
	Name string
	Movie movie.Movie
}

type Config struct {
	Name string
	Movie movie.Movie
}

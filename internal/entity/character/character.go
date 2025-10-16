package character

import (
	"movies/internal/entity/movie"

	"github.com/google/uuid"
)

type Character struct {
	Id uuid.UUID 		`json:"id" validate:"required"`
	Name string			`json:"name" validate:"required"`
	Movie movie.Movie	`json:"movie" validate:"required"`
}

type Option func(movie Character) Character

func WithName(name string) Option {
	return func(character Character) Character {
		c := character
		c.Name = name
		return c
	}
}

func WithMovie(movie movie.Movie) Option {
	return func(character Character) Character {
		c := character
		c.Movie = movie
		return c
	}
}

func NewWithOptions(opts ...Option) *Character {
	character := Character {
		Id: uuid.New(),
	}

	for _, o := range opts {
		character = o(character)
	}

 	return &character
}

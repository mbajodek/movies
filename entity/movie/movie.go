package movie

import (
	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Movie struct {
	Id openapi_types.UUID `json:"id"`
	Year int	 `json:"year"`
	Title string `json:"title"`
}

type Option func(movie Movie) Movie

func WithTitle(title string) Option {
	return func(movie Movie) Movie {
		m := movie
		m.Title = title
		return m
	}
}

func WithYear(year int) Option {
	return func(movie Movie) Movie {
		m := movie
		m.Year = year
		return m
	}
}

func NewWithOptions(opts ...Option) *Movie {
	movie := Movie {
		Id: uuid.New(),
	}

	for _, o := range opts {
		movie = o(movie)
	}

 	return &movie
}

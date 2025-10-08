package movie

import "github.com/google/uuid"

type Movie struct {
	Id uuid.UUID
	Year int
	Title string
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

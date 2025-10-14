package server

import (
	"context"
	"fmt"
	"movies/internal/api"
	"movies/internal/mapper"

	"github.com/go-playground/validator/v10"
)

func (s *Server) PostMovies(_ context.Context, request api.PostMoviesRequestObject) (api.PostMoviesResponseObject, error) {
	validate := validator.New()
	err := validate.Struct(request)

	if err != nil {
		fmt.Println("validation error:", err)
		return nil, err
	}

	title := request.Body.Title
	year := request.Body.Year

	movie := s.Mr.Create(title, year)

	return mapper.MapMovieEntityToPostDto(movie), nil
}

func (s *Server) GetMovies(ctx context.Context, request api.GetMoviesRequestObject) (api.GetMoviesResponseObject, error) {
	movies := s.Mr.GetAll()
	fmt.Println(movies)
	return mapper.MapMovieEntitySliceToDtoSlice(movies), nil
}

func (s *Server) DeleteMoviesId(ctx context.Context, request api.DeleteMoviesIdRequestObject) (api.DeleteMoviesIdResponseObject, error) {
	s.Mr.Delete(request.Id)

	return nil, nil
}

func (s *Server) GetMoviesId(ctx context.Context, request api.GetMoviesIdRequestObject) (api.GetMoviesIdResponseObject, error) {
	movie, exists := s.Mr.Get(request.Id)

	if !exists {
		return nil, nil
	}
	return mapper.MapMovieEntityToGetDto(movie), nil
}

func (s *Server) PutMovies(ctx context.Context, request api.PutMoviesRequestObject) (api.PutMoviesResponseObject, error) {
	validate := validator.New()
	err := validate.Struct(request)

	if err != nil {
		fmt.Println("validation error:", err)
		return nil, err
	}

	id := request.Body.Id
	title := request.Body.Title
	year := request.Body.Year

	movie, err := s.Mr.Update(id, title, year)

	if err != nil {
		fmt.Println("Update movie error:", err)
		return nil, err
	}

	return mapper.MapMovieEntityToUpdateDto(movie), nil
}

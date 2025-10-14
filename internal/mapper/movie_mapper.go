package mapper

import (
	"movies/entity/movie"
	"movies/internal/api"
)

func MapMovieEntityToPostDto(movie movie.Movie) api.PostMovies201JSONResponse {
	return api.PostMovies201JSONResponse(MapMovie(movie))
}

func MapMovieEntitySliceToDtoSlice(movies []movie.Movie) api.GetMovies201JSONResponse {
	res := make(api.GetMovies201JSONResponse, len(movies))
	for i, movie := range movies {
		res[i] = MapMovie(movie)
	}
	return res
}

func MapMovieEntityToGetDto(movie movie.Movie) api.GetMoviesId201JSONResponse {
	return api.GetMoviesId201JSONResponse(MapMovie(movie))
}

func MapMovieEntityToUpdateDto(movie movie.Movie) api.PutMovies201JSONResponse {
	return api.PutMovies201JSONResponse(MapMovie(movie))
}

func MapMovie(movie movie.Movie) api.Movie {
	return api.Movie{
		Id:    movie.Id,
		Title: movie.Title,
		Year:  movie.Year,
	}
}

func MapMovieDtoToEntity(movieDto api.Movie) movie.Movie {
	return movie.Movie{
		Id:    movieDto.Id,
		Title: movieDto.Title,
		Year:  movieDto.Year,
	}
}

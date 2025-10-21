package mapper

import (
	"crypto/rsa"
	"crypto/x509"
	"movies/internal/api"
	"movies/internal/entity/movie"
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

func MapMoviePostRequestBodyToEntity(movieDto *api.PostMoviesJSONRequestBody, cert *x509.Certificate, privateKey *rsa.PrivateKey) movie.Movie {
	return *movie.NewWithOptions(
		movie.WithTitle(movieDto.Title),
		movie.WithYear(movieDto.Year),
		movie.WithCert(cert, privateKey),
	)
}


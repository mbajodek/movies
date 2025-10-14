package server

import (
	"context"
	"fmt"
	"movies/internal/api"
	"movies/internal/mapper"
	"strings"

	"github.com/go-playground/validator/v10"
)

func (s *Server) GetCharacters(ctx context.Context, request api.GetCharactersRequestObject) (api.GetCharactersResponseObject, error) {
	characters := s.Cr.GetAll()

	return mapper.MapCharacterEntitySliceToDtoSlice(characters), nil
}

func (s *Server) GetCharactersId(ctx context.Context, request api.GetCharactersIdRequestObject) (api.GetCharactersIdResponseObject, error) {
	character := s.Cr.Get(request.Id)
	return mapper.MapCharacterEntityToGetDto(character), nil
}

func (s *Server) PostCharacters(ctx context.Context, request api.PostCharactersRequestObject) (api.PostCharactersResponseObject, error) {
	validate := validator.New()
	err := validate.Struct(request)

	if err != nil {
		fmt.Println("validation error:", err)
		return nil, err
	}

	name := request.Body.Name
	movieId := request.Body.MovieId
	movie, exists := s.Mr.Get(movieId)
	
	if !exists {
		msg := fmt.Sprintf("Movie with id: %s does not exist", movieId)
		fmt.Println(msg)
		return api.PostCharacters412JSONResponse{Message: msg}, nil
	}
	
	if isStarWarsMovie(movie.Title) {
		existed, err := s.StarWarsValidator.CharacterExistsInMovie(name)
	
		if !existed {
			msg := fmt.Sprintf("%s does not exists in Star Wars movie", name)
			fmt.Println(msg)
			return api.PostCharacters412JSONResponse{Message: msg}, err
		}
	}

	character := s.Cr.Create(name, movie)

	return mapper.MapCharcterEntityToPostDto(character), nil
}


func isStarWarsMovie(title string) bool {
	return strings.ToLower(title) == "star wars"
}

func (s *Server) PutCharacters(ctx context.Context, request api.PutCharactersRequestObject) (api.PutCharactersResponseObject, error) {
	validate := validator.New()
	err := validate.Struct(request)

	if err != nil {
		fmt.Println("validation error:", err)
		return nil, err
	}

	id := request.Body.Id
	name := request.Body.Name
	movieId := request.Body.MovieId
	movie, exists := s.Mr.Get(*movieId)
	
	if !exists {
		msg := fmt.Sprintf("Movie with id: %s does not exist", movieId)
		fmt.Println(msg)
		return api.PutCharacters412JSONResponse{Message: msg}, nil
	}

	if isStarWarsMovie(movie.Title) {
		existed, err := s.StarWarsValidator.CharacterExistsInMovie(name)

		if !existed {
			msg := fmt.Sprintf("%s does not exists in Star Wars movie", name)
			fmt.Println(msg)
			return api.PutCharacters412JSONResponse{Message: msg}, err
		}
	}

	character, err := s.Cr.Update(id, name, movie)

	if err != nil {
		fmt.Println("Update movie error:", err)
		return nil, err
	}

	return mapper.MapCharacterEntityToUpdateDto(character), nil
}

func (s *Server) DeleteCharactersId(ctx context.Context, request api.DeleteCharactersIdRequestObject) (api.DeleteCharactersIdResponseObject, error) {
	s.Cr.Delete(request.Id)

	return nil, nil
}
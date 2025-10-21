package mapper

import (
	"crypto/rsa"
	"crypto/x509"
	"movies/internal/api"
	"movies/internal/entity/character"
	"movies/internal/entity/movie"
)

func MapCharcterEntityToPostDto(character character.Character) api.PostCharacters201JSONResponse {
	return api.PostCharacters201JSONResponse(MapCharacter(character))
}

func MapCharacterEntitySliceToDtoSlice(characters []character.Character) api.GetCharacters201JSONResponse {
	res := make(api.GetCharacters201JSONResponse, len(characters))
	for i, character := range characters {
		res[i] = MapCharacter(character)
	}
	return res
}

func MapCharacterEntityToGetDto(character character.Character) api.GetCharactersId201JSONResponse {
	return api.GetCharactersId201JSONResponse(MapCharacter(character))
}

func MapCharacterEntityToUpdateDto(character character.Character) api.PutCharacters201JSONResponse {
	return api.PutCharacters201JSONResponse(MapCharacter(character))
}

func MapCharacter(character character.Character) api.Character {
	return api.Character{
		Id:    character.Id,
		Name:  character.Name,
		Movie: MapMovie(character.Movie),
	}
}

func MapCharacterRequestToEntity(name string, movie movie.Movie, cert *x509.Certificate, privateKey *rsa.PrivateKey) character.Character {
	return *character.NewWithOptions(
		character.WithName(name),
		character.WithMovie(movie),
		character.WithCert(cert, privateKey),
	)
}

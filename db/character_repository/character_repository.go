package character_repository

import (
	"errors"
	"fmt"
	"movies/db"
	"movies/entity/character"
	"movies/entity/movie"

	"github.com/google/uuid"
)

type CharacterRepository struct {
	DB *db.MemoryDb
}

func New(db *db.MemoryDb) *CharacterRepository {
	return &CharacterRepository{DB: db}
}

func (mr *CharacterRepository) Create(name string, movie movie.Movie) character.Character {
	character := character.NewWithOptions(
		character.WithName(name),
		character.WithMovie(movie),
	)

	mr.DB.Characters[character.Id] = *character
	return *character
}

func (mr *CharacterRepository) Get(id uuid.UUID) character.Character {
	character, exists :=  mr.DB.Characters[id]
	
	if !exists {
		fmt.Println("No character with id:", id)
		fmt.Println(character)
		return character
	}

	return character
}

func (mr *CharacterRepository) GetAll() []character.Character {
	var characters []character.Character

	for _, character := range mr.DB.Characters {
		characters = append(characters, character)
	}

	return characters
}

func (mr *CharacterRepository) Update(id uuid.UUID, name string, movie movie.Movie) (character.Character, error) {
	character, exists :=  mr.DB.Characters[id]
	
	if !exists {
		fmt.Println("No character with id:", id)
		return character, errors.New("No character with id: " + id.String())
	}
	character.Name = name
	character.Movie = movie
	mr.DB.Characters[id] = character

	return character, nil
}

func (mr *CharacterRepository) Delete(id uuid.UUID) {
	_, exists :=  mr.DB.Characters[id]
	
	if !exists {
		fmt.Println("No character with id:", id)
	}

	delete(mr.DB.Characters, id)
}

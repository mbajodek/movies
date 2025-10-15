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

	mr.DB.Characters.Store(character.Id, *character)
	return *character
}

func (mr *CharacterRepository) Get(id uuid.UUID) character.Character {
	v, ok :=  mr.DB.Characters.Load(id)
	var c character.Character

	if !ok {
		fmt.Println("No character with id:", id)
		return c
	}
	c = v.(character.Character)

	return c
}

func (mr *CharacterRepository) GetAll() []character.Character {
	var characters []character.Character

	mr.DB.Characters.Range(func(key, value interface{}) bool {
        characters = append(characters, value.(character.Character))
        return true
    })

	return characters
}

func (mr *CharacterRepository) Update(id uuid.UUID, name string, movie movie.Movie) (character.Character, error) {
	v, ok :=  mr.DB.Characters.Load(id)
	var c character.Character
	
	if !ok {
		fmt.Println("No character with id:", id)
		return c, errors.New("No character with id: " + id.String())
	}
	c = v.(character.Character)
	c.Name = name
	c.Movie = movie
	mr.DB.Characters.Swap(id, c)

	return c, nil
}

func (mr *CharacterRepository) Delete(id uuid.UUID) {
	_, ok :=  mr.DB.Characters.Load(id)
	
	if !ok {
		fmt.Println("No character with id:", id)
	}

	mr.DB.Characters.Delete(id)
}

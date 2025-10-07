package character

import (
	"fmt"
	"math/rand"
	"slices"
)

func New(cfg Config) Character {
	newCharacter := Character{
		Id:    rand.Int31(),
		Name:  cfg.Name,
		Movie: cfg.Movie,
	}
	Characters = append(Characters, newCharacter)

	return newCharacter
}

func Get(id int32) Character {
	for _, character := range Characters {
		if character.Id == id {
			return character
		}
	}
	fmt.Println("No character with id:", id)
	return Character{}
}

func (character *Character) Update(cfg Config) Character {
	(*character).Name = cfg.Name
	(*character).Movie = cfg.Movie
	return *character
}

func Delete(id int32) {
	for i, character := range Characters {
		if character.Id == id {
			Characters = slices.Delete(Characters, i, i+1)
			break
		}
	}
}
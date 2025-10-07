package main

import(
	"fmt"
	"movies/movie"
	"movies/character"
)

func main() {
	m1 := AddMovie(1994, "Skazani na Shawshank")
	m2 := AddMovie(1999, "Matrix")
	m3 := AddMovie(1999, "Zielona mila")

	c1 := AddCharacter("Morgan Freeman", m1)
	c2 := AddCharacter("Keanu Reeves", m2)

	fmt.Println("Movies:", movie.Movies)
	fmt.Println("Characters:", character.Characters)

	fmt.Println(movie.Get(m1.Id))
	fmt.Println(movie.Get(m2.Id))

	fmt.Println(character.Get(c1.Id))
	fmt.Println(character.Get(c2.Id))

	movie.Movies[1].Update(movie.Config{Year: 1997, Title: "Kiler"})

	character.Characters[0].Update(character.Config{Name: "Tom Hanks", Movie: m3})

	character.Delete(c2.Id)
	fmt.Println("Characters:", character.Characters)

	movie.Delete(m1.Id)
	fmt.Println("Movies:", movie.Movies)
}

func AddMovie(year int, title string) movie.Movie {
	cfg := movie.Config{Year: year, Title: title}
	return movie.New(cfg)
}

func AddCharacter(name string, movie movie.Movie) character.Character {
	cfg := character.Config{Name: name, Movie: movie}
	return character.New(cfg)
}

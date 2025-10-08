package main

import (
	"fmt"
	"movies/db"
	"movies/db/character_repository"
	"movies/db/movie_repository"
)

func main() {
	db := db.New()
	mr := movie_repository.New(db)
	cr := character_repository.New(db)

	m1 := mr.Create("Skazani na Shawshank", 1994)
	m2 := mr.Create("Matrix", 1999)
	m3 := mr.Create("Zielona mila", 1999)

	fmt.Println(m1)

	c1 := cr.Create("Andy", m1,)
	c2 := cr.Create("Morfeus", m2)
	c3 := cr.Create("Neo", m2)

	fmt.Println("Movies:", db.Movies)
	fmt.Println("Characters:", db.Characters)

	fmt.Println(mr.Get(m1.Id))

	fmt.Println(cr.Get(c1.Id))

	mr.Update(m3.Id, "Kiler", 1997)
	cr.Update(c1.Id, "Andy Dufresne", m1)

	cr.Delete(c2.Id)
	cr.Delete(c3.Id)

	mr.Delete(m2.Id)
	fmt.Println("Movies:", db.Movies)
	fmt.Println("Characters:", db.Characters)
}

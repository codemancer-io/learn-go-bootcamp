package main

import "fmt"

type Artist struct {
	Name, Genre string
	Songs       int
}

func newReleasae(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	me := &Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newReleasae(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}

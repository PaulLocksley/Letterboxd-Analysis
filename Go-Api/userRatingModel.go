package main

type userRaiting struct {
	Name   string
	Movies []movie
}

type movie struct {
	Name    string
	Raiting int
	ID      string
	link    string
	Crew    []person
}

type person struct {
	Name string
	Roll string
}

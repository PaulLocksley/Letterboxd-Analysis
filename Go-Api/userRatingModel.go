package main

type userRaiting struct {
	Name   string
	Movies []movie
}

type movie struct {
	Name    string
	Raiting int
	Crew    []person
}

type person struct {
	Name string
	Roll string
}

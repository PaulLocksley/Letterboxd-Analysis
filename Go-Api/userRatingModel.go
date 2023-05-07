package main

type userRaiting struct {
	Name   string
	Movies []movie
}

type movie struct {
	Name    string
	Raiting int
	ID      string
	Link    string
	Crew    []person
}

type person struct {
	Name string
	Role string
}

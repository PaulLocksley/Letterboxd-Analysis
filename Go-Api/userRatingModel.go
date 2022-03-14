package main

type userRaiting struct {
	name   string
	movies []movie
}

type movie struct {
	name    string
	raiting int
	crew    []person
}

type person struct {
	name string
	roll string
}

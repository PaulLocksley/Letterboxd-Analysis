package main

import "fmt"

func main() {
	user := FetchUserRestuls("sigmasalt")
	fmt.Println(user)
	fmt.Println(len(user.movies))
}

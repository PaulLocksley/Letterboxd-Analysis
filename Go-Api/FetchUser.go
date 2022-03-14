package main

//this pulls all raiting data for a user
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func FetchUserRestuls(username string) userRaiting {
	resp, err := http.Get("https://letterboxd.com/" + username + "/films/ratings/")
	userMovies := userRaiting{}
	if err != nil {
		log.Fatalln(err)
	}
	//b, err := goquery.NewDocumentFromReader()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	userMovies = parseList(userMovies, string(b))
	pageCount := getPageCount(string(b))
	for i := 2; i <= pageCount; i++ { //Could add concurrency here for improvmenet in large users.
		extraBody := fetchExtraPage(username, pageCount)
		userMovies = parseList(userMovies, extraBody)
	}
	fmt.Println("Fetch finished, now parsing")

	return userMovies
}

func fetchExtraPage(username string, pageNo int) string {
	resp, err := http.Get("https://letterboxd.com/" + username + "/films/ratings/page/" + strconv.Itoa(pageNo) + "/")
	if err != nil {
		return ""
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(b)
}

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FetchUserRestuls(username string) userRaiting {
	resp, err := http.Get("https://letterboxd.com/" + username + "/films/ratings/")
	userHtmlPages := []*goquery.Document{}
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)
	firstPage, err := goquery.NewDocumentFromReader(strings.NewReader(string(b)))
	userHtmlPages = append(userHtmlPages, firstPage)
	if err != nil {
		log.Fatalln(err)
	}
	pageCount := getPageCount(string(b))
	for i := 2; i <= pageCount; i++ { //TODO: Could add concurrency here for improvmenet in large users.
		userHtmlPages = append(userHtmlPages, fetchExtraPage(username, i))
	}
	fmt.Println("Fetch finished, now parsing")
	return parseList(userHtmlPages, username)
}

func fetchExtraPage(username string, pageNo int) *goquery.Document {
	resp, err := http.Get("https://letterboxd.com/" + username + "/films/ratings/page/" + strconv.Itoa(pageNo) + "/")
	if err != nil {
		log.Fatal(err)
	}
	b, _ := io.ReadAll(resp.Body)
	Page, err := goquery.NewDocumentFromReader(strings.NewReader(string(b)))
	if err != nil {
		log.Fatal(err)
	}
	return Page
}

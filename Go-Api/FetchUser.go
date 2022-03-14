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
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)
	firstPage, err := goquery.NewDocumentFromReader(strings.NewReader(string(b)))
	if err != nil {
		log.Fatalln(err)
	}
	pageCount := getPageCount(string(b))
	userHtmlPages := make([]*goquery.Document, pageCount)
	userHtmlPages[0] = firstPage
	for i := 2; i <= pageCount; i++ {
		//go func(i int, userHTMLPages []*goquery.Document) { //Todo: Work this out so it doesnt take 10 seconds per 500 results
		userHtmlPages[i-1] = fetchExtraPage(username, i)
		//}(i, userHtmlPages)
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

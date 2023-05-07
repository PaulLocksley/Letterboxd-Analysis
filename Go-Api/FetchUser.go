package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

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
	var wg sync.WaitGroup
	for i := 2; i <= pageCount; i++ {
		wg.Add(1)
		go func(i int, userHTMLPages []*goquery.Document) {
			defer wg.Done()
			userHtmlPages[i-1] = fetchExtraPage(username, i)
		}(i, userHtmlPages)
	}
	wg.Wait()
	//fmt.Println("Fetch finished, now parsing")
	return parseList(userHtmlPages, username)
}

func FetchUserRestulsNoDetails(username string) userRaiting {
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
	var wg sync.WaitGroup
	for i := 2; i <= pageCount; i++ {
		wg.Add(1)
		go func(i int, userHTMLPages []*goquery.Document) {
			defer wg.Done()
			userHtmlPages[i-1] = fetchExtraPage(username, i)
		}(i, userHtmlPages)
	}
	wg.Wait()
	//fmt.Println("Fetch finished, now parsing")
	return parseListNoDetails(userHtmlPages, username)
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

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func parseList(userMovieList userRaiting, userData string) userRaiting { //https://zetcode.com/golang/net-html/
	tkn := html.NewTokenizer(strings.NewReader(userData))

	var vals []string

	var isLi bool

	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return userRaiting{}

		case tt == html.StartTagToken:

			t := tkn.Token()
			isLi = t.Data == "ul"
			break
		case tt == html.TextToken:

			t := tkn.Token()

			if isLi {
				vals = append(vals, t.Data)
			}

			isLi = false
		}
	}
	fmt.Println(vals)
	return userRaiting{}
}

func getPageCount(userData string) int {
	r1, err1 := regexp.Compile(`<li class="paginate-page">.*<\/li>`)
	r2, err2 := regexp.Compile(`(\d+)\D*$`)
	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
	}
	match := r1.FindString(userData)
	matchs := r2.FindStringSubmatch(match)
	if len(matchs) < 2 {
		return 0
	}
	number, _ := strconv.Atoi(matchs[len(matchs)-1]) //Because first match and first group match are different.
	return number
}

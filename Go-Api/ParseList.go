package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func parseList(pages []*goquery.Document, username string) userRaiting {
	y := 0
	userRaitings := userRaiting{name: username, movies: []movie{}}
	for i := range pages {
		pages[i].Find(".poster-container").Each(func(x int, s *goquery.Selection) {
			movie := movie{}
			raitingHtml := s.Find(".rating").First()
			raiting, _ := raitingHtml.Html()
			switch raiting {
			case " ½ ":
				movie.raiting = 1
			case " ★ ":
				movie.raiting = 2
			case " ★½ ":
				movie.raiting = 3
			case " ★★ ":
				movie.raiting = 4
			case " ★★½ ":
				movie.raiting = 5
			case " ★★★ ":
				movie.raiting = 6
			case " ★★★½ ":
				movie.raiting = 7
			case " ★★★★ ":
				movie.raiting = 8
			case " ★★★★½ ":
				movie.raiting = 9
			case " ★★★★★ ":
				movie.raiting = 10
			}
			titleHtml := s.Find(".image").First()
			title, _ := titleHtml.Attr("alt")
			movie.name = title
			userRaitings.movies = append(userRaitings.movies, movie)
			y++
		})
	}
	fmt.Println("Pages,", len(pages), "Movie Count", y, " Should be close to ", len(pages)*18)
	return userRaitings
}

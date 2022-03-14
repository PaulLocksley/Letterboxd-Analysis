package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func parseList(pages []*goquery.Document, username string) userRaiting {
	y := 0
	userRaitings := userRaiting{Name: username, Movies: []movie{}}
	for i := range pages {
		pages[i].Find(".poster-container").Each(func(x int, s *goquery.Selection) {
			movie := movie{}
			raitingHtml := s.Find(".rating").First()
			raiting, _ := raitingHtml.Html()
			switch raiting {
			case " ½ ":
				movie.Raiting = 1
			case " ★ ":
				movie.Raiting = 2
			case " ★½ ":
				movie.Raiting = 3
			case " ★★ ":
				movie.Raiting = 4
			case " ★★½ ":
				movie.Raiting = 5
			case " ★★★ ":
				movie.Raiting = 6
			case " ★★★½ ":
				movie.Raiting = 7
			case " ★★★★ ":
				movie.Raiting = 8
			case " ★★★★½ ":
				movie.Raiting = 9
			case " ★★★★★ ":
				movie.Raiting = 10
			}
			titleHtml := s.Find(".image").First()
			title, _ := titleHtml.Attr("alt")
			movie.Name = title
			userRaitings.Movies = append(userRaitings.Movies, movie)
			y++
		})
	}
	fmt.Println("Pages,", len(pages), "Movie Count", y, " Should be close to ", len(pages)*18)
	return userRaitings
}

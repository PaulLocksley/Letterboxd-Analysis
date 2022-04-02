package main

import (
	"fmt"
	"sync"

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
			//fmt.Println(s.Html())
			titleHtml := s.Find(".image").First()
			title, _ := titleHtml.Attr("alt")
			movie.Name = title
			IDHTML := s.Find(".really-lazy-load").First()
			id, _ := IDHTML.Attr("data-film-id")
			movie.ID = id

			link, _ := IDHTML.Attr("data-target-link")
			movie.link = link
			// movie.Crew = parseMovie(link, id)
			userRaitings.Movies = append(userRaitings.Movies, movie)
			y++
		})
	}
	fmt.Println("User ", username, " Pages,", len(pages), "Movie Count", y, " Should be close to ", len(pages)*18)
	if len(pages) > 100 { //Todo: Fix for larger users
		return userRaitings
	}
	//get crew
	var wg sync.WaitGroup
	for i := 0; i < len(userRaitings.Movies); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			userRaitings.Movies[i].Crew = parseMovie(userRaitings.Movies[i].ID, userRaitings.Movies[i].link)
		}(i)
	}
	wg.Wait()

	go writeCacheResults(userRaitings.Movies)
	//userRaitings.Movies[0].Crew = parseMovie(userRaitings.Movies[0].ID, userRaitings.Movies[0].link) //testing mode
	return userRaitings
}

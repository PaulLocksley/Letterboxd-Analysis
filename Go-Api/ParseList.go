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
			case "½":
				movie.Raiting = 1
			case "★":
				movie.Raiting = 2
			case "★½":
				movie.Raiting = 3
			case "★★":
				movie.Raiting = 4
			case "★★½":
				movie.Raiting = 5
			case "★★★":
				movie.Raiting = 6
			case "★★★½":
				movie.Raiting = 7
			case "★★★★":
				movie.Raiting = 8
			case "★★★★½":
				movie.Raiting = 9
			case "★★★★★":
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
			movie.Link = link
			// movie.Crew = parseMovie(link, id)
			if movie.Raiting != 0 {
				userRaitings.Movies = append(userRaitings.Movies, movie)
			}
			//userRaitings.Movies = append(userRaitings.Movies, movie)
			y++
		})
	}
	fmt.Println("User ", username, " Pages,", len(pages), "Movie Count", y, " Should be close to ", len(pages)*18)
	// if len(pages) > 100 { //Todo: Fix for larger users
	// 	for i := 0; i < len(userRaitings.Movies); i++ {
	// 		userRaitings.Movies[i].Crew = parseMovie(userRaitings.Movies[i].ID, userRaitings.Movies[i].link)
	// 	}
	// 	go writeCacheResults(userRaitings.Movies)
	// 	return userRaitings
	// }
	//get crew
	for x := 0; x < len(userRaitings.Movies); x += 20 {
		var wg sync.WaitGroup
		var maxValue = x + 20
		if x+20 > len(userRaitings.Movies) {
			maxValue = len(userRaitings.Movies)
		}
		for i := x; i < maxValue; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				userRaitings.Movies[i].Crew = parseMovie(userRaitings.Movies[i].ID, userRaitings.Movies[i].Link)
			}(i)
		}
		wg.Wait()
	}

	go writeCacheResults(userRaitings.Movies)
	//userRaitings.Movies[0].Crew = parseMovie(userRaitings.Movies[0].ID, userRaitings.Movies[0].link) //testing mode
	return userRaitings
}

func parseListNoDetails(pages []*goquery.Document, username string) userRaiting {
	y := 0
	userRaitings := userRaiting{Name: username, Movies: []movie{}}

	for i := range pages {
		pages[i].Find(".poster-container").Each(func(x int, s *goquery.Selection) {
			movie := movie{}
			raitingHtml := s.Find(".rating").First()
			raiting, _ := raitingHtml.Html()
			switch raiting {
			case "½":
				movie.Raiting = 1
			case "★":
				movie.Raiting = 2
			case "★½":
				movie.Raiting = 3
			case "★★":
				movie.Raiting = 4
			case "★★½":
				movie.Raiting = 5
			case "★★★":
				movie.Raiting = 6
			case "★★★½":
				movie.Raiting = 7
			case "★★★★":
				movie.Raiting = 8
			case "★★★★½":
				movie.Raiting = 9
			case "★★★★★":
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
			movie.Link = link
			// movie.Crew = parseMovie(link, id)
			userRaitings.Movies = append(userRaitings.Movies, movie)
			y++
		})
	}
	fmt.Println("User ", username, " Pages,", len(pages), "Movie Count", y, " Should be close to ", len(pages)*18)

	return userRaitings
}

package main

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//
//tab-genres  tab-details //next week goal
//
func parseMovie(id string, url string) []person {
	if cachedResult, ok := movieCache[id]; ok {
		return cachedResult
	}
	people := []person{}

	resp, err := http.Get("https://letterboxd.com/" + url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	movieDetails, _ := goquery.NewDocumentFromReader(strings.NewReader(string(b)))
	//fmt.Println(movieDetails.Html())
	//average Raiting, Genres and Studios.

	Crew := movieDetails.Find("#tab-crew").First()
	Cast := movieDetails.Find(".cast-list").First()
	Details := movieDetails.Find("#tab-details").First()
	Genre := movieDetails.Find("#tab-genres").First()
	crewText, _ := Crew.Html()
	castText, _ := Cast.Html()
	detailsText, _ := Details.Html()
	genreText, _ := Genre.Html()
	crewList := parseCrew(crewText)
	castList := parseCast(castText)
	detailsList := parseDetails(detailsText)
	genreList := parseGenre(genreText)
	people = append(people, genreList...)
	people = append(people, castList...)
	people = append(people, crewList...)
	people = append(people, detailsList...)

	return people
}

func parseDetails(detailsList string) []person { //TODO: fix this so I am not hiding this in a person so I don't need to rewrite my cache
	details := []person{}
	r1, _ := regexp.Compile(`href="/studio/(\S*)/"`)
	r2, _ := regexp.Compile(`href="/films/(\S*)/"`)

	studioMatchs := r1.FindAllStringSubmatch(detailsList, -1)
	localization := r2.FindAllStringSubmatch(detailsList, -1)
	for i := range studioMatchs {
		crewMember := strings.ReplaceAll(studioMatchs[i][1], "-", " ")
		details = append(details, person{Name: crewMember, Role: "Studio"})
	}

	for i := range localization {
		tmpStrings := strings.Split(localization[i][1], "/")
		details = append(details, person{Name: tmpStrings[1], Role: tmpStrings[0]})

	}
	return details
}

func parseGenre(genreList string) []person {
	genre := []person{}
	r1, _ := regexp.Compile(`href="/films/genre/(\S*)/"`)
	genreMatches := r1.FindAllStringSubmatch(genreList, -1)
	for i := range genreMatches {
		genreName := strings.ReplaceAll(genreMatches[i][1], "-", " ")
		genre = append(genre, person{Name: genreName, Role: "Genre"})
	}
	return genre
}

func parseCast(castList string) []person {
	cast := []person{}
	r1, _ := regexp.Compile(`\>([a-zA-Z]*\s[a-zA-Z]*)</a>`)
	matchs := r1.FindAllStringSubmatch(castList, -1)
	for i := range matchs {
		cast = append(cast, person{Name: matchs[i][1], Role: "Cast"}) //Todo: add nullchecks everywhere.
	}
	return cast
}

func parseCrew(crewText string) []person {
	crew := []person{}
	r1, _ := regexp.Compile(`href="(\S*)"`)
	matchs := r1.FindAllStringSubmatch(crewText, -1)
	for i := range matchs {
		tmpStrings := strings.Split(matchs[i][1], "/")
		crewMember := strings.ReplaceAll(tmpStrings[2], "-", " ")
		crew = append(crew, person{Name: crewMember, Role: tmpStrings[1]})
	}
	return crew
}

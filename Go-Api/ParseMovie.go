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
	people := []person{}

	resp, err := http.Get("https://letterboxd.com/" + url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	movieDetails, _ := goquery.NewDocumentFromReader(strings.NewReader(string(b)))
	Crew := movieDetails.Find("#tab-crew").First()
	Cast := movieDetails.Find(".cast-list").First()
	crewText, _ := Crew.Html()
	castText, _ := Cast.Html()
	crewList := parseCrew(crewText)
	castList := parseCast(castText)
	people = append(people, castList...)
	people = append(people, crewList...)
	return people
}

func parseCast(castList string) []person {
	cast := []person{}
	r1, _ := regexp.Compile(`\>([a-zA-Z]*\s[a-zA-Z]*)</a>`)
	matchs := r1.FindAllStringSubmatch(castList, -1)
	for i := range matchs {
		cast = append(cast, person{Name: matchs[i][1], Roll: "Cast"}) //Todo: add nullchecks everywhere.
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
		crew = append(crew, person{Name: crewMember, Roll: tmpStrings[1]})
	}
	return crew
}

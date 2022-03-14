package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func getPageCount(userData string) int { //TODO: Change this to GoQuery at some point
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

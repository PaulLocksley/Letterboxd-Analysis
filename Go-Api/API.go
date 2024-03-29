package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var movieCache map[string][]person = make(map[string][]person)
var movieCacheLock bool

func main() {
	loadCache()
	fmt.Println("Loaded", len(movieCache), "movies")
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/user", userlist)
	http.HandleFunc("/moviedetails", movieDetails)
	http.HandleFunc("/userNoDetails", userListNoDetails)

	//log.Fatal(http.ListenAndServe(":1313", nil))
	log.Fatal(http.ListenAndServeTLS(":1313", "certificate.crt", "certificate.key", nil))
}

func userListNoDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Referrer Policy", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	args := r.URL.Query()
	user := FetchUserRestulsNoDetails(args.Get("u"))
	JsonUser, _ := json.Marshal(user)
	fmt.Fprint(w, string(JsonUser))
}

func userlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Referrer Policy", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	args := r.URL.Query()
	user := FetchUserRestuls(args.Get("u"))
	JsonUser, _ := json.Marshal(user)
	fmt.Fprint(w, string(JsonUser))
	//fmt.Println("Endpoint Hit: homePage", string(JsonUser), user, "\n")
}

func movieDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Referrer Policy", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	args := r.URL.Query()
	url := args.Get("m")
	movie := movie{Name: url, Raiting: 1, ID: "TEMP", Link: url}
	movie.Crew = parseMovie("TEMP", "film/"+url)
	JsonMovie, _ := json.Marshal(movie)
	fmt.Fprint(w, string(JsonMovie))
}

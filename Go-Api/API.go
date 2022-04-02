package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var movieCache map[string][]person = make(map[string][]person)
var movieCacheLock bool

func userlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Referrer Policy", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	args := r.URL.Query()
	user := FetchUserRestuls(args.Get("u"))
	JsonUser, _ := json.Marshal(user)
	fmt.Fprint(w, string(JsonUser))
	//fmt.Println("Endpoint Hit: homePage", string(JsonUser), user, "\n")
}

func handleRequests() {
	http.HandleFunc("/user", userlist)
	//log.Fatal(http.ListenAndServe(":1313", nil))
	log.Fatal(http.ListenAndServeTLS(":1313", "certificate.crt", "certificate.key", nil))
}

func main() {
	loadCache()
	fmt.Println("Loaded", len(movieCache), "movies")
	handleRequests()
}

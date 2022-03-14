package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func userlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Referrer Policy", "application/json")
	args := r.URL.Query()
	user := FetchUserRestuls(args.Get("u"))
	JsonUser, err := json.Marshal(user)
	fmt.Fprintf(w, string(JsonUser))
	fmt.Println("Endpoint Hit: homePage", string(JsonUser), user, "\n", err)
}

func handleRequests() {
	http.HandleFunc("/user", userlist)
	log.Fatal(http.ListenAndServe(":1313", nil))
}

func main() {
	handleRequests()
}

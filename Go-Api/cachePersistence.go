package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func saveCache() {
	file, err := os.Create("movieCache.txt")
	if err != nil {
		fmt.Println("Save Error", err)
		return
	}

	defer file.Close()
	r, err := Marshal(movieCache)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(file, r)
	if err != nil {
		fmt.Println("Save Error", err)
	}

}

func loadCache() {
	file, err := os.ReadFile("movieCache.txt")
	if err != nil {
		fmt.Println("Load err", err)
	}
	err = json.Unmarshal(file, &movieCache)
	fmt.Println(err)
}

//https://medium.com/@matryer/golang-advent-calendar-day-eleven-persisting-go-objects-to-disk-7caf1ee3d11d
var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

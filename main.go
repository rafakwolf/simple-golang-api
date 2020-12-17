package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article is definition of an article object
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles is an array of Article object
type Articles []Article

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:   "My first golang API",
			Desc:    "Creating some really simple golang API",
			Content: "Golang rocks!",
		},
	}

	fmt.Println("Returring all articles")
	json.NewEncoder(w).Encode(articles)
}

func main() {
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page endpoint")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", getAllArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

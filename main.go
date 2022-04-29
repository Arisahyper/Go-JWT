package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gin-auth/auth"

	"github.com/gorilla/mux"
)

type Post struct {
	ID    int    `json:"id"`
	Tag   string `json:"tag"`
	Title string `json:"title"`
	Body  string `json:"body"`
	URL   string `json:"url"`
}

func main() {
	route := mux.NewRouter()

	route.Handle("/", homeHandler)	// Root
	route.Handle("/private", auth.JwtMiddleware.Handler(private))	// JwtMiddleware check token
	route.Handle("/auth", auth.GetTokenHandler)					// get token

	if err := http.ListenAndServe(":8080", route); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

var homeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
})

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := []Post{
		{ID: 1, Tag: "tag1", Title: "title1", Body: "body1", URL: "url1"},
		{ID: 2, Tag: "tag2", Title: "title2", Body: "body2", URL: "url2"},
		{ID: 3, Tag: "tag3", Title: "title3", Body: "body3", URL: "url3"},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
})

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed"
)

/*
	/feed
	  GET
	  returns a list of feed items
	  encoding/json

	/feed/sources
	  GET
	  returns a list of sources
	  encoding/json

	/feed/sources/add
	  POST
	    body
	      name: string
	      url: string
	  attempt to add given name,url as source
	    check if rss feed can parse correctly
	    if cannot parse
	      fail
	    else
	      returns feed added (w/ id)

	/feed/sources/remove/:id
	  DELETE

	/feed/sources/update/:id
	  UPDATE
	    body
	      name: string
	      url: string

	/feed/sources/:id
	  GET
	    returns feed items from a specific id
*/

func main() {
	r := mux.NewRouter()
	r.Use(handlers.CORS())

	r.HandleFunc("/feed", GetFeed)

	log.Printf("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func GetFeed(w http.ResponseWriter, r *http.Request) {
	url := "https://lobste.rs/rss"
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		log.Fatal("failed to parse rss feed")
	}

	body, err := json.Marshal(feed)
	if err != nil {
		log.Fatal("failed to marshall feed into json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

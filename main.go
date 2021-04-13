package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed"
)

type Source struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

var sources = map[uint64]*Source{
	1: {1, "lobsters", "https://lobste.rs/rss"},
	2: {2, "shtetl-optimized", "https://www.scottaaronson.com/blog/?feed=rss2"},
	3: {3, "hackernews", "https://news.ycombinator.com/rss"},
}

func main() {
	r := mux.NewRouter()
	r.Use(handlers.CORS())

	r.HandleFunc("/sources", GetSources)
	r.HandleFunc("/sources/add", AddSource).Methods("POST")
	r.HandleFunc("/sources/del/{id}", RemoveSource).Methods("POST")
	r.HandleFunc("/sources/edit/{id}", EditSource).Methods("POST")
	r.HandleFunc("/feed/{id}", GetFeed)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("client/public/")))

	log.Printf("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func GetSources(w http.ResponseWriter, r *http.Request) {
	var response []*Source

	for _, value := range sources {
		response = append(response, value)
	}

	body, err := json.Marshal(response)
	if err != nil {
		log.Fatal("failed to marshal sources into json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func AddSource(w http.ResponseWriter, r *http.Request) {
	var source Source
	err := json.NewDecoder(r.Body).Decode(&source)
	if err != nil {
		log.Fatal("bad request body!")
	}

	sources[0] = &source
}

func RemoveSource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)

	delete(sources, id)
}

func EditSource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)

	var nameurl struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&nameurl)
	if err != nil {
		log.Fatal("bad request body")
	}

	sources[id].Name = nameurl.Name
	sources[id].Url = nameurl.Url
}

func GetFeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	source := sources[id]

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(source.Url)
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

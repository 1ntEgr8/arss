package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed"
	"gorm.io/gorm"
)

type SourceHandler struct {
	db *gorm.DB
}

func NewSourceHandler(db *gorm.DB) *SourceHandler {
	return &SourceHandler{
		db: db,
	}
}

func (s *SourceHandler) GetSources(w http.ResponseWriter, r *http.Request) {
	var sources []Source
	s.db.Find(&sources)

	body, err := json.Marshal(sources)
	if err != nil {
		log.Fatal("failed to marshal sources into json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func (s *SourceHandler) AddSource(w http.ResponseWriter, r *http.Request) {
	var source Source
	err := json.NewDecoder(r.Body).Decode(&source)
	if err != nil {
		log.Fatal("bad request body!")
	}
	s.db.Create(&source)

	body, err := json.Marshal(source)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func (s *SourceHandler) RemoveSource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s.db.Delete(&Source{}, vars["id"])
}

func (s *SourceHandler) EditSource(w http.ResponseWriter, r *http.Request) {
	var source Source

	nameUrl := struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}{}
	vars := mux.Vars(r)
	s.db.Find(&source, vars["id"])

	err := json.NewDecoder(r.Body).Decode(&nameUrl)
	if err != nil {
		log.Fatal("bad request body")
	}

	source.Name = nameUrl.Name
	source.Url = nameUrl.Url

	s.db.Save(&source)
}

func (s *SourceHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	var source Source

	vars := mux.Vars(r)
	s.db.First(&source, vars["id"])

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

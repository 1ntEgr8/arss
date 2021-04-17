package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed"
	"gorm.io/gorm"
)

type Error struct {
	Msg string `json:"msg"`
}

type SourceHandler struct {
	db *gorm.DB
}

func NewSourceHandler(db *gorm.DB) *SourceHandler {
	return &SourceHandler{
		db: db,
	}
}

func (s *SourceHandler) GetSources(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var sources []Source
	s.db.Find(&sources)

	body, err := json.Marshal(sources)
	if err != nil {
		msg := "unable to grab sources"
		log.Println(msg)
		json.NewEncoder(w).Encode(Error{msg})
	} else {
		w.Write(body)
	}
}

func (s *SourceHandler) AddSource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var source Source
	err := json.NewDecoder(r.Body).Decode(&source)
	if err != nil {
		msg := "failed to decode request body"
		log.Println(msg)
		json.NewEncoder(w).Encode(Error{msg})
		return
	}
	s.db.Create(&source)
	body, err := json.Marshal(source)
	w.Write(body)
}

func (s *SourceHandler) RemoveSource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s.db.Delete(&Source{}, vars["id"])
}

func (s *SourceHandler) EditSource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var source Source

	nameUrl := struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}{}
	vars := mux.Vars(r)
	s.db.Find(&source, vars["id"])

	err := json.NewDecoder(r.Body).Decode(&nameUrl)
	if err != nil {
		msg := "failed to decode request body"
		log.Println(msg)
		json.NewEncoder(w).Encode(Error{msg})
		return
	}

	source.Name = nameUrl.Name
	source.Url = nameUrl.Url

	s.db.Save(&source)
}

func (s *SourceHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var source Source
	vars := mux.Vars(r)
	s.db.First(&source, vars["id"])

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(source.Url)
	if err != nil {
		msg := "failed to fetch feed; perhaps the URL is invalid"
		log.Println(msg)
		json.NewEncoder(w).Encode(Error{msg})
		return
	}

	body, err := json.Marshal(feed)
	if err != nil {
		msg := "failed to build feed"
		log.Println(msg)
		json.NewEncoder(w).Encode(Error{msg})
		return
	}

	w.Write(body)
}

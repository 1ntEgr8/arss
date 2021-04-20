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

func Json(v interface{}, w http.ResponseWriter) error {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	return nil
}

func InternalServerError(msg string, w http.ResponseWriter) {
	log.Println(msg)
	w.WriteHeader(http.StatusInternalServerError)
	Json(Error{msg}, w)
}

func (s *SourceHandler) GetSources(w http.ResponseWriter, r *http.Request) {
	var sources []Source
	s.db.Find(&sources)

	err := Json(sources, w)
	if err != nil {
		InternalServerError("failed to jsonify sources", w)
	}
}

func (s *SourceHandler) AddSource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var source Source
	err := json.NewDecoder(r.Body).Decode(&source)
	if err != nil {
		InternalServerError("failed to decode body", w)
		return
	}
	s.db.Create(&source)
	Json(source, w)
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
		InternalServerError("failed to decode request body", w)
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
		InternalServerError("failed to fetch feed; perhaps the URL is invalid", w)
		return
	}

	err = Json(feed, w)
	if err != nil {
		InternalServerError("failed to make feed", w)
		return
	}
}

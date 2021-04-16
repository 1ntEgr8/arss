package main

import (
	"log"
	"net/http"
	_ "strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mmcdole/gofeed"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Source struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	Serve()
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sources.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Source{})
	return db
}

func Serve() {
	db := ConnectDB()
	s := NewSourceHandler(db)

	r := mux.NewRouter()
	r.Use(handlers.CORS())

	r.HandleFunc("/sources", s.GetSources)
	r.HandleFunc("/sources/add", s.AddSource).Methods("POST")
	r.HandleFunc("/sources/del/{id}", s.RemoveSource).Methods("POST")
	r.HandleFunc("/sources/edit/{id}", s.EditSource).Methods("POST")
	r.HandleFunc("/feed/{id}", s.GetFeed)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("client/public/")))

	log.Printf("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

package main

import (
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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

	log.Printf("Listening on port 8080")
	go open("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// https://stackoverflow.com/questions/39320371/how-start-web-server-to-open-page-in-browser-in-golang
// open opens the specified URL in the default browser of the user.
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

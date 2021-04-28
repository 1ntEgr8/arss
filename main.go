package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	headless := flag.Bool("headless", false, "run server in headless mode")
	port := flag.Int("port", 8080, "which port to run server on")
	client := flag.String("client-path", "client/public", "path to client")
	dbPath := flag.Bool("db-path", false, "prints path to sources db")
	flag.Parse()

	if *dbPath {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(fmt.Sprintf("%s/sources.db", pwd))
	} else {
		Serve(*client, *port, *headless)
	}
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sources.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Source{})
	return db
}

func Serve(client string, port int, headless bool) {
	db := ConnectDB()
	s := NewSourceHandler(db)

	r := mux.NewRouter()
	r.Use(handlers.CORS())

	r.HandleFunc("/sources", s.GetSources)
	r.HandleFunc("/sources/add", s.AddSource).Methods("POST")
	r.HandleFunc("/sources/del/{id}", s.RemoveSource).Methods("POST")
	r.HandleFunc("/sources/edit/{id}", s.EditSource).Methods("POST")
	r.HandleFunc("/feed/{id}", s.GetFeed)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(client)))

	if !headless {
		go open(fmt.Sprintf("http://localhost:%d", port))
	} else {
		log.Printf("Running in headless mode")
	}
	log.Printf(fmt.Sprintf("Listening on port %d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
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

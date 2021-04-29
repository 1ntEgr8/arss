package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not determine user's home directory")
	}

	defaultClientPath := filepath.Join(homeDir, ".arss", "clients", "default")

	headless := flag.Bool("headless", false, "run server in headless mode")
	port := flag.Int("port", 8080, "which port to run server on")
	clientPath := flag.String("client", defaultClientPath, "path to client")
	printConfigPath := flag.Bool("config-path", false, "prints path to default configuration files")
	flag.Parse()

	if *printConfigPath {
		configPath := filepath.Join(homeDir, ".arss")
		fmt.Printf(configPath)
	} else {
		Serve(*clientPath, *port, *headless)
	}
}

func ConnectDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Source{})
	return db
}

func Serve(clientPath string, port int, headless bool) {
	clientExists, err := exists(clientPath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not check if client at %s exists", clientPath))
	}
	if !clientExists {
		log.Fatal(fmt.Sprintf("%s does not exist", clientPath))
	}

	db := ConnectDB(filepath.Join(clientPath, "sources.db"))
	s := NewSourceHandler(db)

	r := mux.NewRouter()
	r.Use(handlers.CORS())

	r.HandleFunc("/sources", s.GetSources)
	r.HandleFunc("/sources/add", s.AddSource).Methods("POST")
	r.HandleFunc("/sources/del/{id}", s.RemoveSource).Methods("POST")
	r.HandleFunc("/sources/edit/{id}", s.EditSource).Methods("POST")
	r.HandleFunc("/feed/{id}", s.GetFeed)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(clientPath)))

	if !headless {
		go open(fmt.Sprintf("http://localhost:%d", port))
	} else {
		log.Printf("Running in headless mode")
	}
	log.Printf(fmt.Sprintf("Listening on port %d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

// helpers

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

// https://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

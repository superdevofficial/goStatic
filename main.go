package main

import (
	"flag"
	"log"
	"net/http"
	"time"
    "fmt"
	"github.com/gorilla/mux"
    "os"
)

var (
        dir string
        defaultFile string
)

func main() {

    flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
    flag.Parse()
    defaultFile = dir + "/index.html"

    r := mux.NewRouter()

    r.PathPrefix("/").HandlerFunc(StaticHandler)
    r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	fmt.Printf("Starting golang server\n")

    srv := &http.Server{
        Handler:      r,
        Addr:         "0.0.0.0:8080",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("File not found %s",r.URL)
    http.ServeFile(w, r, defaultFile)
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	path := dir + r.URL.EscapedPath()

	if _, err := os.Stat(path); err == nil {
		http.ServeFile(w, r, path)
	} else {
		http.ServeFile(w, r, defaultFile)
	}
}

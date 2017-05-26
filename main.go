package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	flag.Parse()
	root := "./"        // 1st argument is the directory location
	extention := ".brs" // 2nd argument is the file extention
	walker := NewWalker(root, walkDirectory(extention))
	walker.Walk()

	router := NewRouter()

	server := &http.Server{
		Addr:           determineListenAddress(),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}

func determineListenAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":8080"
	}
	return ":" + port
}

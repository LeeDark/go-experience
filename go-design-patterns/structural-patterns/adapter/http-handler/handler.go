package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyServer struct {
	Msg string
}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.Msg)
}

func main() {
	server := &MyServer{
		Msg: "Hello, World",
	}

	http.Handle("/", server)

	// The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.
	http.HandleFunc("/func", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, func!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

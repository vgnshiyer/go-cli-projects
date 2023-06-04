package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(rw, "ParseForm err: %v", err)
		return
	}
	fmt.Fprintf(rw, "POST request succesful\n")
	name := r.FormValue("name")
	fmt.Fprintf(rw, "Hello %s\n", name)
}

func helloHandler(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(rw, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(rw, "method is not acceptable", http.StatusNotFound)
		return
	}
	fmt.Fprintf(rw, "Hello!")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Serving at localhost:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
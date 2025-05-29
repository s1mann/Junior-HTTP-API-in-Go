package main

import (
	"encoding/json"
	f "fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

var people []Person

func main() {
	http.HandleFunc("/people", peopleHandler)
	http.HandleFunc("/health", healthCheckHandler)

	log.Println("server start listeng on port 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPeople(w)
	case http.MethodPost:
		postPerson(w, r)
	default:
		http.Error(w, "ivailer http method", http.StatusMethodNotAllowed)
	}
}

func getPeople(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(people)
	f.Fprintf(w, "get people: '%v'", people)
}

func postPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	people = append(people, person)
	f.Fprintf(w, "post new person: '%v'", person)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	f.Fprint(w, "http web-server works correctly")
}

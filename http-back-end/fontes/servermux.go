package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Home\n")
}

// função principal para executar a api
func main() {
	fmt.Println("Iniciando servermux...")
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Paul", Lastname: "Silva", Address: &Address{City: "City W", State: "State W"}})
	router.HandleFunc("/contato", GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}/{name}", CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")
	router.HandleFunc("/", home).Methods("GET")
	log.Fatal(http.ListenAndServe(":8090", router))
}

package main

import (
	"boardgame-tracker/src/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/person", controllers.CreatePerson)
	log.Fatal(http.ListenAndServe(":8080", router))
}

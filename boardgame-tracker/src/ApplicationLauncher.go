package main

import (
	"boardgame-tracker/src/pkg/controllers"
	"boardgame-tracker/src/pkg/utility"
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
	test := utility.GetHotGames()
	print(test)
	log.Fatal(http.ListenAndServe(":8080", router))

}

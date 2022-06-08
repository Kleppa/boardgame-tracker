package controllers

import (
	"boardgame-tracker/src/pkg/documents"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getBoardGameList(w http.ResponseWriter, r *http.Request) {

}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var newPerson documents.Person
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println( "wrong payload")
	}
	json.Unmarshal(reqBody, &newPerson)
	w.WriteHeader(http.StatusCreated)
	encoderErr := json.NewEncoder(w).Encode(newPerson)
	if encoderErr != nil {
		return
	}
	fmt.Printf("%p person created\n", newPerson)

}

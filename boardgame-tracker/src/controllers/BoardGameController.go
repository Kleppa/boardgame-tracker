package controllers

import (
	"boardgame-tracker/src/documents"
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
		fmt.Fprintf(w, "wrong payload")
	}
	json.Unmarshal(reqBody, &newPerson)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(newPerson)
	fmt.Fprintf(w, "%p person created", newPerson)

}

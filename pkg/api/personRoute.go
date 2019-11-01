package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/QOSQOs/UNIVeasier/pkg/app/personService"
	"github.com/gorilla/mux"
)

func (s *Server) addPerson(w http.ResponseWriter, r *http.Request) {
	err := personService.AddPerson(s.Conn, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := []byte("The record was added successfully")

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Server) getPersonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	person, err := personService.GetPersonById(s.Conn, int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Server) getListPerson(w http.ResponseWriter, r *http.Request) {
	personList, err := personService.GetListPerson(s.Conn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(personList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/QOSQOs/UNIVeasier/pkg/app/interestService"
	"github.com/gorilla/mux"
)

func (s *Server) addInterest(w http.ResponseWriter, r *http.Request) {
	err := interestService.AddInterest(s.Conn, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := []byte("The record was added successfully")

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Server) getInterestById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	interest, err := interestService.GetInterestById(s.Conn, int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(&interest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Server) getListInterest(w http.ResponseWriter, r *http.Request) {
	interestList, err := interestService.GetListInterest(s.Conn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(interestList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

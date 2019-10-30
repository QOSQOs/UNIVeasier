package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/QOSQOs/UNIVeasier/pkg/app/typeUniversityService"
	"github.com/gorilla/mux"
)

func (s *Server) addTypeUniversity(w http.ResponseWriter, r *http.Request) {
	err := typeUniversityService.AddTypeUniversity(s.Conn, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := []byte("The record was successfully added")

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Server) getTypeUniversityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	record, err := typeUniversityService.GetTypeUniversityById(s.Conn, int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Server) getListTypeUniversity(w http.ResponseWriter, r *http.Request) {
	records, err := typeUniversityService.GetListTypeUniversity(s.Conn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// should we delete a register or only disable it?
func (s *Server) deleteTypeUniversityById(w http.ResponseWriter, r *http.Request) {
	return
}

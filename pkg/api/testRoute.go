package api

import (
	"encoding/json"
	"net/http"

	"github.com/QOSQOs/UNIVeasier/pkg/app/testService"
	"github.com/gorilla/mux"
)

func (s *Server) getTest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	test, err := testService.GetTest(s.Conn, vars["name"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(test)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Server) getListTest(w http.ResponseWriter, r *http.Request) {
	tests, err := testService.GetListTest(s.Conn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(tests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Server) addTest(w http.ResponseWriter, r *http.Request) {
	err := testService.AddTest(s.Conn, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := []byte("The record was successfully added")

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

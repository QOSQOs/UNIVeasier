package api

import (
	"encoding/json"
	"net/http"

	"github.com/QOSQOs/UNIVeasier/pkg/app"
	"github.com/gorilla/mux"
)

func (s *Server) RoutesTest() {
	s.Router.HandleFunc("/test/{name}", s.getTest).Methods("GET")
	s.Router.HandleFunc("/test", s.getListTest).Methods("GET")
}

func (s *Server) getTest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	test, err := app.ServiceGetTest(s.Conn, vars["name"])
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
	tests, err := app.ServiceGetListTest(s.Conn)
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

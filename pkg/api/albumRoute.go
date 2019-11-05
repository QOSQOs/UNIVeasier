package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/QOSQOs/UNIVeasier/pkg/app/albumService"
	"github.com/gorilla/mux"
)

func (s *Server) addAlbum(w http.ResponseWriter, r *http.Request) {
	err := albumService.AddAlbum(s.Conn, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := []byte("The record was added successfully")

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Server) getAlbumById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	albumModel, err := albumService.GetAlbumById(s.Conn, int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(&albumModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Server) getListAlbum(w http.ResponseWriter, r *http.Request) {
	albumList, err := albumService.GetListAlbum(s.Conn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(albumList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

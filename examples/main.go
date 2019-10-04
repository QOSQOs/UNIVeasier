package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Test struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Server struct {
	Router *mux.Router
	db     *sql.DB
}

func (s *Server) initialize() error {
	// Open up our database connection
	db, err := sql.Open("mysql", "root:qosqo123@tcp(localhost:3306)/dbtest")
	if err != nil {
		fmt.Println("Error the arguments are not valid:", err.Error())
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connection with the database:", err.Error())
		return err
	}

	fmt.Println("Successfully connected")
	s.db = db
	s.Router = mux.NewRouter()
	s.initializeRoutes()

	return nil
}

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/test/{name}", s.getTest).Methods("GET")
	s.Router.HandleFunc("/test", s.listTests).Methods("GET")
}

func (s *Server) run() error {
	fmt.Println("Listening on http://localhost:8000")

	err := http.ListenAndServe(":8000", s.Router)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	return s.db.Close()
}

func (s *Server) getTest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	test, err := get(s.db, vars["name"])
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

func (s *Server) listTests(w http.ResponseWriter, r *http.Request) {
	tests, err := list(s.db)
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

func get(db *sql.DB, name string) (Test, error) {
	// Execute query to return one registry
	var test Test

	err := db.QueryRow("SELECT id, name FROM dbtest.tabletest where name = ?", name).Scan(&test.ID, &test.Name)
	if err != nil {
		return Test{}, err
	}

	return test, nil
}

func list(db *sql.DB) ([]Test, error) {
	// Execute query to return many registries
	results, err := db.Query("SELECT id, name FROM dbtest.tabletest")
	if err != nil {
		panic(err.Error())
	}

	var listTest []Test
	for results.Next() {
		var test Test
		err = results.Scan(&test.ID, &test.Name)
		if err != nil {
			return nil, err
		}
		listTest = append(listTest, test)
	}

	return listTest, nil
}

func main() {
	s := Server{}
	s.initialize()
	s.run()
}

package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"go.uber.org/zap"
)

var log *zap.SugaredLogger

type Test struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Server struct {
	Router     *mux.Router
	db         *sql.DB
	configPath string
	config     *Config
}

func (s *Server) initialize() error {
	// Open up our database connection
	config, err := GetConfig(s.configPath)
	if err != nil {
		return err
	}

	db, err := InitConnection(config.DB)
	if err != nil {
		return err
	}

	log.Info("Successfully connected")

	s.config = config
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

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Info("Listening on http://localhost:8000")

	err := http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(s.Router))
	if err != nil {
		log.Errorw("Error listening", "error", err.Error())
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
	logger, _ := zap.NewDevelopment()
	log = logger.Sugar()

	configFilepath := flag.String("config", "", "References to configuration file.")
	flag.Parse()

	log.Info(*configFilepath)

	s := Server{configPath: *configFilepath}
	s.initialize()
	s.run()
}

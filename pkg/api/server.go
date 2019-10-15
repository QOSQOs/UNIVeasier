package api

import (
	"database/sql"
	"net/http"

	"github.com/QOSQOs/UNIVeasier/internal/common"
	"github.com/QOSQOs/UNIVeasier/internal/config"
	"github.com/QOSQOs/UNIVeasier/pkg/db/connection"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Server struct {
	Log    *zap.SugaredLogger
	Config *config.Config
	Conn   *sql.DB
	Router *mux.Router
}

func (s *Server) initLogger() error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	s.Log = logger.Sugar()
	return nil
}

func (s *Server) initConfiguration(filename string) error {
	var err error
	s.Config, err = config.ReadConfiguration(filename)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) initDBConnection() error {
	var err error
	s.Conn, err = connection.InitConnection(s.Config)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) initRouter() {
	s.Router = mux.NewRouter()
	s.RoutesTest()
}

func (s *Server) Initialize(filename string) error {
	err := s.initLogger()
	if err != nil {
		return err
	}
	common.Log = s.Log
	common.Global = "variable"

	s.Log.Info("Reading the configuration file")
	err = s.initConfiguration(filename)
	if err != nil {
		return err
	}

	s.Log.Info("Initializing database connection")
	err = s.initDBConnection()
	if err != nil {
		return err
	}

	s.Log.Info("Configuring the router")
	s.initRouter()

	return nil
}

func (s *Server) Run() {
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	s.Log.Info("Server listening on http://localhost:8000")

	err := http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(s.Router))
	if err != nil {
		s.Log.Errorw("Cannot initialize server", "info", err.Error())
	} else {
		s.Conn.Close()
	}
}

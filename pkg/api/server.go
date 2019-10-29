package api

import (
	"database/sql"
	"fmt"
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
	Config *config.Config
	Conn   *sql.DB
	Router *mux.Router
}

func (s *Server) initLogger() error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	common.Log = logger.Sugar()
	return nil
}

func (s *Server) initConfiguration(filepath string) error {
	var err error
	s.Config, err = config.ReadConfiguration(filepath)
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

	common.Log.Info("Reading the configuration file")
	err = s.initConfiguration(filename)
	if err != nil {
		return err
	}

	common.Log.Info("Initializing database connection")
	err = s.initDBConnection()
	if err != nil {
		return err
	}

	common.Log.Info("Configuring the router")
	s.initRouter()

	return nil
}

func (s *Server) Run() {
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	common.Log.Info(fmt.Sprintf("Server listening on http://%s:%s", s.Config.SERVER.HOST, s.Config.SERVER.PORT))

	port := fmt.Sprintf(":%s", s.Config.SERVER.PORT)
	err := http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(s.Router))
	if err != nil {
		common.Log.Errorw("Cannot initialize server", "info", err.Error())
	} else {
		s.Conn.Close()
	}
}

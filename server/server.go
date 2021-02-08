package server

import (
	"database/sql"
	"fmt"
	"net/http"

	//IMplicit import
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rameesThattarath/qaForum/handler"
	"github.com/rameesThattarath/qaForum/repository"
	"github.com/rameesThattarath/qaForum/usecase"

	"log"
)

//Server - Server
type Server struct {
	Router      *mux.Router
	UserHandler *handler.UserHandler
}

//Initialize server
func (s *Server) Initialize(config *Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := sql.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := repository.NewSQLUserRepo(db)
	h := &handler.UserHandler{
		RegisterUC: &usecase.RegisterUserImpl{
			Repo: userRepo,
		},
		LoginUC: &usecase.RegisterUserImpl{
			Repo: userRepo,
		},
		ProfileUC: &usecase.GetProfileImpl{
			Repo: userRepo,
		},
	}

	s.UserHandler = h
	s.Router = mux.NewRouter()
	s.setRouters()
}

// Run the app on it's router
func (s *Server) Run(host string) {
	log.Fatal(http.ListenAndServe(host, s.Router))
}

// Config - Config
type Config struct {
	DB *dBConfig
}

type dBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

//GetConfig - GetConfig
func GetConfig() *Config {
	return &Config{
		DB: &dBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "rameesdb",
			Name:     "qa_forum",
			Charset:  "utf8",
		},
	}
}

func (s *Server) setRouters() {
	s.Router.HandleFunc("/", s.home()).Methods("GET")
	s.Router.HandleFunc("/signup", s.signup()).Methods("POST")
	s.Router.HandleFunc("/login", s.login()).Methods("POST")
	s.Router.HandleFunc("/profile", isAuthorized(s.profile())).Methods("GET")
}

func (s *Server) home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "I aint dead")
	}
}

func (s *Server) signup() http.HandlerFunc {
	return s.UserHandler.HandleSignup()
}

func (s *Server) login() http.HandlerFunc {
	return s.UserHandler.HandleLogin()
}

func (s *Server) profile() http.HandlerFunc {
	return s.UserHandler.HandleProfile()
}

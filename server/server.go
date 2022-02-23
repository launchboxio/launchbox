package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robwittman/launchbox/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServerOpts struct {
	Port int64
}

type Server struct {
	r *gin.Engine
}

var database *gorm.DB

func New(opts *ServerOpts) *Server {
	r := gin.Default()
	server := &Server{r: r}

	initServer()

	server.initControllers()

	return server
}

func initServer() {
	fmt.Println("Initing our server")
	db, err := gorm.Open(postgres.Open("host=localhost user=launchbox password=password dbname=launchbox port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(
		&api.Application{},
		&api.Project{},
		&api.Revision{},
	)
	if err != nil {
		panic(err)
	}
	database = db
}

func (s *Server) Run() error {
	return s.r.Run()
}

func (s *Server) initControllers() {
	(&Applications{}).Register(s.r)
	(&Projects{}).Register(s.r)
	(&Revisions{}).Register(s.r)
}

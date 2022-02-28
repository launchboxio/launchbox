package server

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robwittman/launchbox/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type ServerOpts struct {
	Port     int64
	RedisUrl string
}

type Server struct {
	r *gin.Engine
}

var database *gorm.DB
var taskServer *machinery.Server

func Run(opts *ServerOpts) error {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	server := &Server{r: r}

	initServer()
	ts, err := NewTaskServer(&TaskServerConfig{RedisUrl: opts.RedisUrl})
	if err != nil {
		panic(err)
	}
	taskServer = ts
	go func() {
		err := RunWorker("machinery_tasks")
		if err != nil {
			panic(err)
		}
	}()
	server.initControllers()

	err = server.Run()
	return err
}

func initServer() {
	db, err := gorm.Open(postgres.Open("host=localhost user=launchbox password=password dbname=launchbox port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(
		&api.Application{},
		&api.Project{},
		&api.Revision{},
		&api.Secret{},
		&api.Build{},
		&api.Task{},
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

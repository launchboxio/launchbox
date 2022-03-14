package server

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/launchboxio/launchbox/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	r *gin.Engine
}

var database *gorm.DB
var taskServer *machinery.Server

func Run(configFile string) error {
	// First things first load our config
	config, err := LoadDefaultConfig(configFile)
	if err != nil {
		return err
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     config.Cors.AllowOrigins,
		AllowMethods:     config.Cors.AllowMethods,
		AllowHeaders:     config.Cors.AllowHeaders,
		ExposeHeaders:    config.Cors.ExposeHeaders,
		AllowCredentials: config.Cors.AllowCredentials,
		MaxAge:           config.Cors.MaxAge,
	}))
	server := &Server{r: r}

	initServer()
	ts, err := NewTaskServer(&TaskServerConfig{
		RedisUrl: config.Redis.Url,
	})

	if err != nil {
		panic(err)
	}
	taskServer = ts
	go func() {
		err := RunWorker(config.Worker.ConsumerTag)
		if err != nil {
			panic(err)
		}
	}()
	server.initControllers(config)

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

func (s *Server) initControllers(config *Config) {
	(&Applications{}).Register(s.r)
	(&Projects{}).Register(s.r)
	(&Revisions{}).Register(s.r)
	(&Logs{}).Register(s.r)
	(&Metrics{
		Config: config.Prometheus,
	}).Register(s.r)
}

package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/launchboxio/launchbox/api"
	"github.com/launchboxio/launchbox/internal/config"
	"github.com/launchboxio/launchbox/internal/resources/applications"
	"github.com/launchboxio/launchbox/internal/resources/metrics"
	"github.com/launchboxio/launchbox/internal/resources/projects"
	"github.com/launchboxio/launchbox/internal/resources/revisions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Config
	r      *gin.Engine
}

var database *gorm.DB

func New(configFile string) (*Server, error) {
	// First things first load our config
	conf, err := config.LoadDefaultConfig(configFile)
	if err != nil {
		return nil, err
	}

	server := &Server{
		config: conf,
	}

	server.init()
	server.registerRoutes()

	return server, nil
}

func (s *Server) init() {
	db, err := gorm.Open(postgres.Open(s.config.Database.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(
		&api.Account{},
		&api.Application{},
		&api.Project{},
		&api.Revision{},
		&api.Secret{},
		&api.Build{},
		&api.Session{},
		&api.Task{},
		&api.User{},
		&api.VerificationToken{},
		&api.Webhook{},
	)
	if err != nil {
		panic(err)
	}
	database = db
}

func (s *Server) Run() error {
	return s.r.Run()
}

func (s *Server) registerRoutes() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     s.config.Cors.AllowOrigins,
		AllowMethods:     s.config.Cors.AllowMethods,
		AllowHeaders:     s.config.Cors.AllowHeaders,
		ExposeHeaders:    s.config.Cors.ExposeHeaders,
		AllowCredentials: s.config.Cors.AllowCredentials,
		MaxAge:           s.config.Cors.MaxAge,
	}))

	apps := applications.New(database)
	group := r.Group("/applications")
	{
		group.GET("", apps.List)
		group.GET("/:applicationId", apps.Get)
		group.POST("", apps.Create)
		group.PUT("/:applicationId", apps.Update)
		group.DELETE("/:applicationId", apps.Delete)
	}

	proj := projects.New(database)
	revisionsCtrl := revisions.New(database)

	projectsGroup := r.Group("/projects")
	{
		projectsGroup.GET("", proj.List)
		projectsGroup.GET("/:projectId", proj.Get)
		projectsGroup.POST("", proj.Create)
		projectsGroup.PUT("/:projectId", proj.Update)
		projectsGroup.DELETE("/:projectId", proj.Delete)

		revisionsGroup := projectsGroup.Group("/:projectId/revisions")
		{
			revisionsGroup.GET("", revisionsCtrl.List)
			revisionsGroup.GET("/:revisionId", revisionsCtrl.Get)
			revisionsGroup.POST("", revisionsCtrl.Create)
			revisionsGroup.PUT("/:revisionId", revisionsCtrl.Update)
			revisionsGroup.DELETE("/:revisionId", revisionsCtrl.Delete)
		}

		webhookCtrl := &Webhooks{}
		webhooksGroup := projectsGroup.Group("/:projectId/webhooks")
		{
			webhooksGroup.GET("", webhookCtrl.List)
			webhooksGroup.GET("/:webhookId", webhookCtrl.Get)
			webhooksGroup.POST("", webhookCtrl.Create)
			webhooksGroup.POST("/:webhookToken", webhookCtrl.Receive)
			webhooksGroup.PUT("/:webhookId", webhookCtrl.Update)
			webhooksGroup.DELETE("/:webhookId", webhookCtrl.Delete)
		}

	}

	metricsCtrl := metrics.New(s.config.Prometheus)
	metricsGroup := r.Group("/metrics")
	//metricsGroup.Use(verifyMetricsEnabled(metricsCtrl.config))
	metricsGroup.GET("", metricsCtrl.Query)

	s.r = r
}

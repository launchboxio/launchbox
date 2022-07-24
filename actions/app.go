package actions

import (
	"net/http"
	"os"
	"strings"

	"github.com/gobuffalo/logger"
	"github.com/sirupsen/logrus"

	"launchbox/locales"
	"launchbox/models"
	"launchbox/public"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/v3/pop/popmw"
	"github.com/gobuffalo/envy"
	csrf "github.com/gobuffalo/mw-csrf"
	forcessl "github.com/gobuffalo/mw-forcessl"
	i18n "github.com/gobuffalo/mw-i18n/v2"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/unrolled/secure"

	gwa "github.com/gobuffalo/gocraft-work-adapter"
	"github.com/gomodule/redigo/redis"

	"github.com/markbates/goth/gothic"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")

var (
	app *buffalo.App
	T   *i18n.Translator
)

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_launchbox_session",
			Worker: gwa.New(gwa.Options{
				Pool: &redis.Pool{
					MaxActive: 5,
					MaxIdle:   5,
					Wait:      true,
					Dial: func() (redis.Conn, error) {
						return redis.Dial("tcp", ":6379")
					},
				},
				Name:           "launchbox",
				MaxConcurrency: 25,
			}),

			Logger: JSONLogger(getLogLevel(envy.Get("LOG_LEVEL", "debug"))),
		})

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//   c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))
		// Setup and use translations:
		app.Use(translations())

		app.GET("/", HomeHandler)

		//AuthMiddlewares
		app.Use(SetCurrentUser)
		app.Use(Authorize)

		////Routes for Auth
		//auth := app.Group("/auth")
		//auth.GET("/", AuthLanding)
		//auth.GET("/new", AuthNew)
		//auth.POST("/", AuthCreate)
		//auth.DELETE("/", AuthDestroy)
		//auth.Middleware.Skip(Authorize, AuthLanding, AuthNew, AuthCreate)

		//Routes for User registration
		users := app.Group("/users")
		users.GET("/new", UsersNew)
		users.POST("/", UsersCreate)
		users.Middleware.Remove(Authorize)

		apps := ApplicationsResource{}
		wr := app.Resource("/applications", apps)
		wr.Resource("projects", ProjectsResource{}).Middleware.Use(SetCurrentApplication)
		wr.Resource("revisions", RevisionsResource{}).Middleware.Use(SetCurrentApplication)

		wr.Middleware.Use(SetCurrentApplication)
		wr.Middleware.Skip(SetCurrentApplication, apps.List, apps.Create)

		app.Resource("/clusters", ClustersResource{})

		auth := app.Group("/auth")
		auth.GET("/{provider}", buffalo.WrapHandlerFunc(gothic.BeginAuthHandler))
		auth.GET("/{provider}/callback", AuthCallback)
		app.ServeFiles("/", http.FS(public.FS())) // serve files from the public directory
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(locales.FS(), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}

func JSONLogger(lvl logger.Level) logger.FieldLogger {
	l := logrus.New()
	l.Level = lvl
	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetOutput(os.Stdout)
	return logger.Logrus{FieldLogger: l}
}

func getLogLevel(logLevel string) logger.Level {
	switch strings.ToLower(logLevel) {
	case "error":
		return logger.ErrorLevel
	case "fatal":
		return logger.FatalLevel
	case "debug":
		return logger.DebugLevel
	case "warn":
		return logger.WarnLevel
	default:
		return logger.InfoLevel
	}
}

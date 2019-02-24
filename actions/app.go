package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/cpjudge/cpjudge_webserver/models"
	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
			SessionName: "_cpjudge_webserver_session",
		})
		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Set the request content type to JSON
		app.Use(middleware.SetContentType("application/json"))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		authenticateGroup := app.Group("/")
		authenticateGroup.Use(AuthenticationMiddleware)
		authenticateGroup.GET("/", HomeHandler)
		app.GET("/websocket", WebSocketHandler)
		app.POST("/users", SignupHandler)
		authenticateGroup.GET("/users/{username}", GetUserInfoHandler)
		app.POST("/login", SigninHandler).Name("login")

		authenticateGroup.POST("/question", QuestionHandler)
		authenticateGroup.GET("/questions", GetQuestionsHandler)
		authenticateGroup.GET("/question/{question_id}", GetQuestionHandler)

		authenticateGroup.POST("/contest", ContestHandler)
		authenticateGroup.GET("/contest", GetContestsHandler)
		authenticateGroup.GET("/contest/{contest_id}", GetContestHandler)

		authenticateGroup.POST("/participate_in", ParticipateInHandler)
		uthenticateGroup.GET("/participates_in/{user_id}", GetParticipatesInHandler)

	}

	return app
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return ssl.ForceSSL(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}

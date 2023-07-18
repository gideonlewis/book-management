package http

import (
	"net/http"
	"regexp"

	echoSentry "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"git.teqnological.asia/teq-go/teq-echo/config"
	"git.teqnological.asia/teq-go/teq-echo/delivery/http/book"
	"git.teqnological.asia/teq-go/teq-echo/delivery/http/borrow"
	"git.teqnological.asia/teq-go/teq-echo/delivery/http/example"
	"git.teqnological.asia/teq-go/teq-echo/delivery/http/healthcheck"
	"git.teqnological.asia/teq-go/teq-echo/delivery/http/user"
	"git.teqnological.asia/teq-go/teq-echo/usecase"
)

func NewHTTPHandler(useCase *usecase.UseCase) *echo.Echo {
	var (
		e         = echo.New()
		loggerCfg = middleware.DefaultLoggerConfig
	)

	loggerCfg.Skipper = func(c echo.Context) bool {
		return c.Request().URL.Path == "/health-check"
	}

	e.Use(middleware.LoggerWithConfig(loggerCfg))
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(echoSentry.New(echoSentry.Options{Repanic: true}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOriginFunc: func(origin string) (bool, error) {
			return regexp.MatchString(
				`^https:\/\/(|[a-zA-Z0-9]+[a-zA-Z0-9-._]*[a-zA-Z0-9]+\.)teqnological.asia$`,
				origin,
			)
		},
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch,
			http.MethodPost, http.MethodDelete, http.MethodOptions,
		},
	}))
	//ping - pong
	e.GET("/ping", echoSwagger.WrapHandler)

	// API docs
	if !config.GetConfig().Stage.IsProd() {
		e.GET("/docs/*", echoSwagger.WrapHandler)
	}

	// test swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Health check
	healthcheck.Init(e.Group("/health-check"))

	// APIs
	api := e.Group("/api")
	api.Use(authMiddleware)
	example.Init(api.Group("/examples"), useCase)
	user.Init(api.Group("/users"), useCase)
	book.Init(api.Group("/books"), useCase)
	borrow.Init(api.Group("/borrows"), useCase)
	return e
}

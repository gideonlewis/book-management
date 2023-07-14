package main

import (
	"context"
	"flag"
	"net"
	"time"

	"git.teqnological.asia/teq-go/teq-pkg/teqlogger"
	"git.teqnological.asia/teq-go/teq-pkg/teqsentry"
	sentryGo "github.com/getsentry/sentry-go"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"git.teqnological.asia/teq-go/teq-echo/client/mysql"
	"git.teqnological.asia/teq-go/teq-echo/config"
	serviceGRPC "git.teqnological.asia/teq-go/teq-echo/delivery/grpc"
	serviceHttp "git.teqnological.asia/teq-go/teq-echo/delivery/http"
	"git.teqnological.asia/teq-go/teq-echo/delivery/job"
	"git.teqnological.asia/teq-go/teq-echo/docs"
	"git.teqnological.asia/teq-go/teq-echo/migration"
	"git.teqnological.asia/teq-go/teq-echo/proto"
	"git.teqnological.asia/teq-go/teq-echo/repository"
	"git.teqnological.asia/teq-go/teq-echo/usecase"
)

const VERSION = "1.0.0"

// @title Example API
// @version 1.0

// @BasePath /api
// @schemes http https

// @securityDefinitions.apikey AuthToken
// @in header
// @name Authorization

// @description Transaction API.
func main() {
	var (
		cfg     = config.GetConfig()
		taskPtr = flag.String("task", "server", "server")
	)

	flag.Parse()

	// sentry
	if len(cfg.SentryDSN) > 0 {
		err := sentryGo.Init(sentryGo.ClientOptions{
			Dsn:              cfg.SentryDSN,
			Environment:      cfg.Stage.ToString(),
			AttachStacktrace: true,
			Release:          VERSION,
		})
		if err != nil {
			teqlogger.GetLogger().Fatal(err.Error())
		}
	}

	// setup locale
	{
		loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
		if err != nil {
			teqsentry.Fatal(err)
			teqlogger.GetLogger().Fatal(err.Error())
		}

		time.Local = loc
	}

	client := mysql.GetClient
	repo := repository.New(client)
	useCase := usecase.New(repo)

	switch *taskPtr {
	case "server":
		executeServer(useCase, client)
	default:
		executeServer(useCase, client)
	}
}

func executeServer(useCase *usecase.UseCase, client func(ctx context.Context) *gorm.DB) {
	cfg := config.GetConfig()

	// migration
	migration.Up(client(context.Background()))

	// cronjob
	if len(cfg.HealthCheck.HealthCheckEndPoint) > 0 {
		job.New().Run()
	}

	// swagger
	docs.SwaggerInfo.Host = cfg.ServiceHost

	l, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}

	m := cmux.New(l)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())
	errs := make(chan error)

	// http
	{
		h := serviceHttp.NewHTTPHandler(useCase)
		go func() {
			h.Listener = httpL
			errs <- h.Start("")
		}()
	}

	// gRPC
	{
		s := grpc.NewServer()

		grpcServer := &serviceGRPC.TeqService{UseCase: useCase}
		proto.RegisterTeqServiceServer(s, grpcServer)

		go func() {
			errs <- s.Serve(grpcL)
		}()
	}

	go func() {
		errs <- m.Serve()
	}()

	err = <-errs
	if err != nil {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}
}

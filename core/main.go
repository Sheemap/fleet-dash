package main

import (
	"fleet-dash-core/data"
	"fleet-dash-core/endpoints"
	pb "fleet-dash-core/protobuf"
	"fleet-dash-core/service"
	"fleet-dash-core/transports"
	"fleet-dash-core/utilities"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	connectionString, ok := os.LookupEnv("COCKROACH_DB_CONNECTION_STRING")
	if !ok {
		_ = level.Error(logger).Log("message", "COCKROACH_DB_CONNECTION_STRING environment variable not set")
		os.Exit(1)
	}

	repository := data.NewRepository(connectionString)

	eventIngestionService := service.NewEventIngestionService(repository)
	eventStreamService := service.NewEventStreamService(repository, logger)
	sessionService := service.NewSessionService(repository, logger)
	staticDataService := service.NewStaticDataService(repository, logger)

	eventEndpoints := endpoints.MakeGrpcEndpoints(eventIngestionService, sessionService, logger)
	httpEndpoints := endpoints.MakeHttpEndpoints(logger, sessionService, eventStreamService, staticDataService)

	jwtValidator := utilities.NewValidator()
	grpcServer := transports.NewGRPCServer(eventEndpoints, logger, jwtValidator)
	httpServer := transports.MakeHTTPHandler(httpEndpoints, logger, jwtValidator)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		_ = logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	// Use cmux to liston for both HTTP and gRPC.
	m := cmux.New(listener)

	grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))

	httpListener := m.Match(cmux.Any())

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterFleetDashServiceServer(baseServer, grpcServer)
		_ = level.Info(logger).Log("transport", "gRPC", "addr", ":8080")
		errs <- baseServer.Serve(grpcListener)
	}()

	go func() {
		_ = logger.Log("transport", "HTTP", "addr", ":8080")
		errs <- http.Serve(httpListener, httpServer)
	}()

	errs <- m.Serve()

	_ = level.Error(logger).Log("exit", <-errs)
}

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
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//type server struct{}
//
//func (s *server) PostEveLogEvent(ctx context.Context, event *protobuf.EveLogEvent) (*protobuf.EveLogEventResponse, error) {
//	fmt.Println("Received event: ", event)
//	return &protobuf.EveLogEventResponse{}, nil
//}

func main() {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)



	repository := data.NewRepository()

	eventService := service.NewEventService(repository)
	sessionService := service.NewSessionService(repository, logger)

	eventEndpoints := endpoints.MakeGrpcEndpoints(eventService, sessionService)
	httpEndpoints := endpoints.MakeHttpEndpoints(sessionService)

	jwtValidator := utilities.NewValidator()
	grpcServer := transports.NewGRPCServer(eventEndpoints, logger, jwtValidator)
	httpServer := transports.MakeHTTPHandler(httpEndpoints, logger, jwtValidator)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterFleetDashServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("transport", "gRPC", "addr", ":50051")
		errs <- baseServer.Serve(grpcListener)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", ":8080")
		errs <- http.ListenAndServe(":8080", httpServer)
	}()

	level.Error(logger).Log("exit", <-errs)

	//db := data.Init()
	//
	//listener, err := net.Listen("tcp", "0.0.0.0:50051")
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//fmt.Println("grpc listening :50051")
	//
	//s := grpc.NewServer()
	//protobuf.RegisterFleetDashServiceServer(s, &server{})
	//
	//go func() {
	//	err := s.Serve(listener)
	//	if err != nil {
	//		log.Fatalf("failed to serve: %v", err)
	//	}
	//}()
	//
	//r := mux.NewRouter()
	//
	//apiRouter := r.PathPrefix("/api").Subrouter()
	//api.RegisterRoutes(apiRouter, db)
	//
	//fmt.Println("https listening :18443")
	//err = http.ListenAndServeTLS(":18443", "ssl/ca.crt", "ssl/ca.key", r)
	//fmt.Println(err)
}

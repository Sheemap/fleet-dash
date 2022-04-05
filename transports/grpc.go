package transports

import (
	"context"
	"fleet-dash-core/endpoints"
	pb "fleet-dash-core/protobuf"
	"fleet-dash-core/utilities"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	gt "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

type gRPCServer struct {
	pb.UnimplementedFleetDashServiceServer

	postEveLogEventBatch gt.Handler
}

func NewGRPCServer(endpoints endpoints.GrpcEndpoints, logger log.Logger, validator utilities.JwtValidator) pb.FleetDashServiceServer {
	errHandler := &grpcErrorHandler{logger: logger}
	options := []gt.ServerOption{
		gt.ServerBefore(extractGrpcAuthentication(validator)),
		gt.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		gt.ServerErrorHandler(errHandler),
	}
	return &gRPCServer{
		postEveLogEventBatch: gt.NewServer(
			endpoints.PostEveLogEventBatch,
			decodePostEveLogEventRequest,
			encodeEveLogEventResponse,
			options...,
		),
	}
}

func (s *gRPCServer) PostEveLogEventBatch(ctx context.Context, req *pb.EveLogEventBatch) (*pb.BatchedEveLogEventResponse, error) {
	_, _, err := s.postEveLogEventBatch.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.BatchedEveLogEventResponse{}, nil
}

func decodePostEveLogEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.EveLogEventBatch)
	return req, nil
}

func encodeEveLogEventResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return &pb.BatchedEveLogEventResponse{}, nil
}

func extractGrpcAuthentication(validator utilities.JwtValidator) func(context.Context, metadata.MD) context.Context {
	return func(ctx context.Context, md metadata.MD) context.Context {
		tokenMeta := md.Get("Authorization")
		if len(tokenMeta) == 0 {
			return ctx
		}
		if tokenMeta[0] == "" {
			return ctx
		}

		splitToken := strings.Split(tokenMeta[0], " ")
		if len(splitToken) != 2 {
			return ctx
		}

		tokenType := splitToken[0]
		token := splitToken[1]

		if tokenType != "Bearer" {
			return ctx
		}

		parsed, err := validator.Parse(token)
		if err != nil || !parsed.Valid {
			return ctx
		}

		charID, err := utilities.GetCharacterIDFromToken(parsed)
		if err != nil {
			return ctx
		}
		name, err := utilities.GetCharacterNameFromToken(parsed)
		if err != nil {
			return ctx
		}

		ctx = context.WithValue(ctx, "character_id", charID)
		ctx = context.WithValue(ctx, "character_name", name)
		return context.WithValue(ctx, "character_access_token", parsed)
	}
}

type grpcErrorHandler struct {
	logger log.Logger
}

func (eh *grpcErrorHandler) Handle(_ context.Context, _ error) {

}

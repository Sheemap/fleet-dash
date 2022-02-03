package endpoints

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	pb "fleet-dash-core/protobuf"
	"fleet-dash-core/service"
	"github.com/go-kit/kit/endpoint"
)

type GrpcEndpoints struct {
	PostEveLogEvent endpoint.Endpoint
}

type EmptyResponse struct {}

func MakeGrpcEndpoints(es service.EventIngestionService, ss service.SessionService, l log.Logger) GrpcEndpoints {

	postEveLogEventEndpoint := makePostEveLogEventEndpoint(es, ss)
	postEveLogEventEndpoint = requireAuthenticated(postEveLogEventEndpoint)
	postEveLogEventEndpoint = setGrpcErrorCodes(postEveLogEventEndpoint)
	postEveLogEventEndpoint = grpcLoggingMiddleware(postEveLogEventEndpoint, l)


	return GrpcEndpoints{
		PostEveLogEvent: postEveLogEventEndpoint,
	}
}

func makePostEveLogEventEndpoint(es service.EventIngestionService, ss service.SessionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		token := ctx.Value("character_access_token")
		if token == nil {
			return nil, ErrNotAuthenticated
		}

		// We require an active session before ingesting any events
		parsed := token.(*jwt.Token)
		sessionId, err := ss.GetCharacterActiveSession(parsed)
		if err != nil {
			return nil, err
		}
		if sessionId == nil {
			return nil, service.ErrNotInSession
		}

		req := request.(*pb.EveLogEvent)

		event := service.EveLogEvent{
			Type: req.Event,
			CharacterID: req.CharacterId,
			Timestamp: req.Timestamp.AsTime(),
			Amount: req.Amount,
			Pilot: req.Pilot,
			Ship: req.Ship,
			Weapon: req.Weapon,
			Application: req.Application,
			Corporation: req.Corporation,
			Alliance: req.Alliance,
		}
		err = es.PersistEveLogEvent(*sessionId, event)
		if err != nil{
			return nil, err
		}

		return EmptyResponse{}, nil
	}
}

func grpcLoggingMiddleware(next endpoint.Endpoint, l log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		method, _ := grpc.Method(ctx)
		defer func(begin time.Time) {
			l.Log(
				"endpoint", method,
				"took", time.Since(begin),
			)
		}(time.Now())
		return next(ctx, request)
	}
}

func setGrpcErrorCodes(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		response, err = next(ctx, request)
		if err != nil {
			return nil, status.Error(codeFromErr(err), err.Error())
		}
		return response, nil
	}
}

func codeFromErr(err error) codes.Code {
	switch err {
	case ErrNotAuthenticated:
		return codes.Unauthenticated
	case service.ErrNotInSession:
		return codes.FailedPrecondition
	default:
		return codes.Internal
	}
}
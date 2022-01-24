package endpoints

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "fleet-dash-core/protobuf"
	"fleet-dash-core/service"
	"github.com/go-kit/kit/endpoint"
)

type GrpcEndpoints struct {
	PostEveLogEvent endpoint.Endpoint
}

type EmptyResponse struct {}

var (
	ErrNotInSession = errors.New("not in active session")
)

func MakeGrpcEndpoints(es service.EventService, ss service.SessionService) GrpcEndpoints {
	return GrpcEndpoints{
		PostEveLogEvent: setGrpcErrorCodes(requireAuthenticated(makePostEveLogEventEndpoint(es, ss))),
	}
}

func makePostEveLogEventEndpoint(es service.EventService, ss service.SessionService) endpoint.Endpoint {
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
			return nil, ErrNotInSession
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
	case ErrNotInSession:
		return codes.FailedPrecondition
	default:
		return codes.Internal
	}
}
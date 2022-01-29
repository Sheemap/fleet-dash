package endpoints

import (
	"context"
	"errors"
	"fleet-dash-core/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	ErrNotAuthenticated = errors.New("not authenticated")
)

type HttpEndpoints struct{
	StartSession endpoint.Endpoint
}

type SessionStartResponse struct {
	SessionId string `json:"sessionId"`
}

func MakeHttpEndpoints(s service.SessionService, l log.Logger) HttpEndpoints {

	startSessionEndpoint := makeStartSessionEndpoint(s)
	startSessionEndpoint = requireAuthenticated(startSessionEndpoint)
	startSessionEndpoint = loggingMiddleware(startSessionEndpoint, l, "startSession")

	return HttpEndpoints{
		StartSession: startSessionEndpoint,
	}
}

func loggingMiddleware(next endpoint.Endpoint, l log.Logger, method string) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		defer func(begin time.Time) {
			l.Log(
				"method", method,
				"took", time.Since(begin),
			)
		}(time.Now())
		return next(ctx, request)
	}
}

func makeStartSessionEndpoint(s service.SessionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		character, err := getCharacter(ctx)
		if err != nil {
			return nil, err
		}

		sessionId, err := s.StartSession(*character)
		if err != nil{
			return nil, err
		}

		return &SessionStartResponse{
			SessionId: sessionId,
		}, nil
	}
}

func getCharacter(ctx context.Context) (*service.Character, error){
	charID := ctx.Value("character_id")
	charName := ctx.Value("character_name")
	charToken := ctx.Value("character_access_token")

	if charID == nil {
		return nil, ErrNotAuthenticated
	}
	if charName == nil {
		return nil, ErrNotAuthenticated
	}
	if charToken == nil {
		return nil, ErrNotAuthenticated
	}


	return &service.Character{
		ID: charID.(string),
		Name: charName.(string),
		AccessToken: charToken.(*jwt.Token),
	}, nil
}
package endpoints

import (
	"context"
	"errors"
	"fleet-dash-core/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang-jwt/jwt/v4"
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

func MakeHttpEndpoints(s service.SessionService) HttpEndpoints {
	return HttpEndpoints{
		StartSession: requireAuthenticated(makeStartSessionEndpoint(s)),
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
package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func requireAuthenticated(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		accessToken := ctx.Value("character_access_token")
		if accessToken == nil {
			return nil, ErrNotAuthenticated
		}

		return next(ctx, request)
	}
}
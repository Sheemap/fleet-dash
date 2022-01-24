package transports

import (
	"context"
	"encoding/json"
	"fleet-dash-core/endpoints"
	"fleet-dash-core/service"
	"fleet-dash-core/utilities"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func MakeHTTPHandler(e endpoints.HttpEndpoints, logger log.Logger, validator utilities.JwtValidator) http.Handler {
	r := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerBefore(extractHttpAuthentication(validator)),
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	api := r.PathPrefix("/api").Subrouter()


	api.Methods("POST").Path("/session/start").Handler(httptransport.NewServer(
		e.StartSession,
		decodeEmptyRequest,
		encodeResponse,
		options...,
	))
	return r
}

func decodeEmptyRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func extractHttpAuthentication(validator utilities.JwtValidator) func(ctx context.Context, r *http.Request) context.Context {
	return func(ctx context.Context, r *http.Request) context.Context {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			return ctx
		}

		splitToken := strings.Split(tokenHeader, " ")
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

// errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error. For more information, read the
// big comment in endpoints.go.
type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case service.ErrSessionAlreadyRunning, service.ErrNotInFleet:
		return http.StatusBadRequest
	case endpoints.ErrNotAuthenticated:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
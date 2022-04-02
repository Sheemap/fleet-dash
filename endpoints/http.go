package endpoints

import (
	"context"
	"errors"
	"fleet-dash-core/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	ErrNotAuthenticated = errors.New("not authenticated")
	ErrIdIsRequired     = errors.New("id is required in query string")
	ErrInvalidId        = errors.New("invalid id")
)

type HttpEndpoints struct {
	EventStreamTicket     endpoint.Endpoint
	EventStream           func(http.ResponseWriter, *http.Request)
	StartSession          endpoint.Endpoint
	EndSession            endpoint.Endpoint
	StaticItemInfo        endpoint.Endpoint
	StaticSolarSystemInfo endpoint.Endpoint
}

type SessionStartResponse struct {
	SessionId string `json:"sessionId"`
}

type GenerateTicketResponse struct {
	Ticket string `json:"ticket"`
}

type NoContentResponse struct{}

func MakeHttpEndpoints(l log.Logger, s service.SessionService, es service.EventStreamService, sds service.StaticDataService) HttpEndpoints {
	streamTicketEndpoint := makeEventStreamTicketEndpoint(s, es)
	streamTicketEndpoint = requireAuthenticated(streamTicketEndpoint)
	streamTicketEndpoint = loggingMiddleware(streamTicketEndpoint, l, "generateStreamTicket")

	eventStream := makeEventStream(es)

	startSessionEndpoint := makeStartSessionEndpoint(s)
	startSessionEndpoint = requireAuthenticated(startSessionEndpoint)
	startSessionEndpoint = loggingMiddleware(startSessionEndpoint, l, "startSession")

	endSessionEndpoint := makeEndSessionEndpoint(s)
	endSessionEndpoint = requireAuthenticated(endSessionEndpoint)
	endSessionEndpoint = loggingMiddleware(endSessionEndpoint, l, "endSession")

	staticItemInfoEndpoint := makeStaticItemInfoEndpoint(sds)
	staticItemInfoEndpoint = requireAuthenticated(staticItemInfoEndpoint)
	staticItemInfoEndpoint = loggingMiddleware(staticItemInfoEndpoint, l, "staticItemInfo")

	staticSolarInfoEndpoint := makeStaticSolarSystemInfoEndpoint(sds)
	staticSolarInfoEndpoint = requireAuthenticated(staticSolarInfoEndpoint)
	staticSolarInfoEndpoint = loggingMiddleware(staticSolarInfoEndpoint, l, "staticSolarSystemInfo")

	return HttpEndpoints{
		EventStreamTicket:     streamTicketEndpoint,
		EventStream:           eventStream,
		StartSession:          startSessionEndpoint,
		EndSession:            endSessionEndpoint,
		StaticItemInfo:        staticItemInfoEndpoint,
		StaticSolarSystemInfo: staticSolarInfoEndpoint,
	}
}

func loggingMiddleware(next endpoint.Endpoint, l log.Logger, method string) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		defer func(begin time.Time) {
			l.Log(
				"endpoint", method,
				"took", time.Since(begin),
			)
		}(time.Now())
		return next(ctx, request)
	}
}

func makeStaticItemInfoEndpoint(s service.StaticDataService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(url.Values)

		if !req.Has("id") {
			return nil, ErrIdIsRequired
		}

		id := req.Get("id")

		idInt, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return nil, ErrInvalidId
		}

		return s.GetStaticItem(idInt)
	}
}

func makeStaticSolarSystemInfoEndpoint(s service.StaticDataService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(url.Values)

		if !req.Has("id") {
			return nil, ErrIdIsRequired
		}

		id := req.Get("id")

		idInt, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return nil, ErrInvalidId
		}

		return s.GetStaticSolarSystem(idInt)
	}
}

func makeEventStreamTicketEndpoint(s service.SessionService, es service.EventStreamService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		character, err := getCharacter(ctx)
		if err != nil {
			return nil, err
		}

		sessionId, err := s.GetCharacterActiveSession(character.AccessToken)
		if err != nil {
			return nil, err
		}
		if sessionId == nil {
			return nil, service.ErrNotInSession
		}

		ticket, err := es.GenerateEventStreamTicket(*sessionId)
		if err != nil {
			return nil, err
		}

		return &GenerateTicketResponse{
			Ticket: *ticket,
		}, nil
	}
}

func makeEventStream(es service.EventStreamService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		ticket := r.URL.Query().Get("ticket")
		activeTicket, err := es.GetActiveTicket(ticket)
		if err != nil {
			// []byte is skipping 2 characters for some reason, so we pad it with 2 spaces
			_ = conn.WriteControl(websocket.CloseMessage, []byte("  invalid ticket"), time.Now().Add(time.Second))
			time.Sleep(time.Second)
			_ = conn.Close()
			return
		}
		if activeTicket == nil {
			_ = conn.WriteControl(websocket.CloseMessage, []byte("  invalid ticket"), time.Now().Add(time.Second))
			time.Sleep(time.Second)
			_ = conn.Close()
			return
		}

		err = es.RegisterStream(activeTicket.SessionID, conn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func makeStartSessionEndpoint(s service.SessionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		character, err := getCharacter(ctx)
		if err != nil {
			return nil, err
		}

		sessionId, err := s.StartSession(*character)
		if err != nil {
			return nil, err
		}

		return &SessionStartResponse{
			SessionId: sessionId,
		}, nil
	}
}

func makeEndSessionEndpoint(s service.SessionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		character, err := getCharacter(ctx)
		if err != nil {
			return nil, err
		}

		err = s.EndSession(*character)
		if err != nil {
			return nil, err
		}

		return &NoContentResponse{}, nil
	}
}

func getCharacter(ctx context.Context) (*service.Character, error) {
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
		ID:          charID.(string),
		Name:        charName.(string),
		AccessToken: charToken.(*jwt.Token),
	}, nil
}

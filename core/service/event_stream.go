package service

import (
	"fleet-dash-core/data"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/websocket"
	"time"
)

type stream struct {
	sessionID string
	ticketID  string
	conn      *websocket.Conn
}

type eventStreamService struct {
	logger          log.Logger
	repo            data.Repository
	streamMap       map[string]map[string]*stream
	events          chan *[]data.Event
	fetchedDbEvents map[string]time.Time
}

type EventStreamService interface {
	GenerateEventStreamTicket(sessionId string) (*string, error)
	GetActiveTicket(ticket string) (*data.EventStreamTicket, error)
	RegisterStream(activeTicket *data.EventStreamTicket, conn *websocket.Conn) error
}

func NewEventStreamService(r data.Repository, l log.Logger) EventStreamService {
	svc := &eventStreamService{
		logger:          l,
		repo:            r,
		streamMap:       make(map[string]map[string]*stream),
		events:          make(chan *[]data.Event, 100),
		fetchedDbEvents: make(map[string]time.Time),
	}

	go svc.startDBPoll()
	go svc.startEventStream()

	return svc
}

func (e *eventStreamService) GenerateEventStreamTicket(sessionId string) (*string, error) {
	ticket, err := e.repo.GenerateEventStreamTicket(sessionId)
	if err != nil {
		return nil, err
	}

	return &ticket.ID, nil
}

func (e *eventStreamService) GetActiveTicket(ticket string) (*data.EventStreamTicket, error) {
	activeTicket, err := e.repo.GetActiveTicket(ticket)
	if err != nil {
		return nil, err
	}

	return activeTicket, nil
}

func (e *eventStreamService) RegisterStream(ticket *data.EventStreamTicket, conn *websocket.Conn) error {
	// If streamMap doesn't have the sessionId, create a new map
	if _, ok := e.streamMap[ticket.SessionID]; !ok {
		e.streamMap[ticket.SessionID] = make(map[string]*stream)
	}
	e.streamMap[ticket.SessionID][ticket.ID] = &stream{
		sessionID: ticket.SessionID,
		ticketID:  ticket.ID,
		conn:      conn,
	}

	// We don't receive close events unless we actively listen for them
	go func(sessionId string, ticketId string, c *websocket.Conn) {
		for {
			if _, _, err := c.NextReader(); err != nil {
				// Remove their conn from the map
				delete(e.streamMap[sessionId], ticketId)
				if len(e.streamMap[sessionId]) == 0 {
					delete(e.streamMap, sessionId)
				}
				_ = c.Close()
				break
			}
		}
	}(ticket.SessionID, ticket.ID, conn)
	return nil
}

func (e *eventStreamService) startEventStream() {
	for {
		events := <-e.events
		if len(*events) == 0 || len(e.streamMap) == 0 {
			continue
		}

		// Group events by session ID
		eventsBySession := make(map[string][]data.Event)
		for _, event := range *events {
			eventsBySession[event.SessionID] = append(eventsBySession[event.SessionID], event)
		}

		// Send events to streams
		for sessionId, streams := range e.streamMap {
			if _, ok := eventsBySession[sessionId]; !ok {
				continue
			}

			for _, stream := range streams {
				err := stream.conn.WriteJSON(eventsBySession[sessionId])
				if err != nil {
					// Remove errored stream
					delete(e.streamMap[sessionId], stream.ticketID)
					if len(e.streamMap[sessionId]) == 0 {
						delete(e.streamMap, sessionId)
					}
				}
			}
		}
	}
}

var SessionClearInterval = time.Minute * 5
var LastSessionClear = time.Now().Add(-SessionClearInterval)
var MapClearInterval = time.Minute * 5
var LastMapClear = time.Now()

func (e *eventStreamService) startDBPoll() {
	for {
		events, err := e.repo.GetEvents(time.Now().Add(-time.Second * 3))
		if err != nil {
			continue
		}

		relevantEvents := e.filterEvents(*events)
		if len(relevantEvents) > 0 {
			e.events <- &relevantEvents
			for _, event := range relevantEvents {
				e.fetchedDbEvents[event.ID] = time.Now()
			}
		}

		if time.Now().Sub(LastSessionClear) > SessionClearInterval {
			// Kick it off in a goroutine so we don't block the main loop
			go e.clearOldSessions()
		}

		if time.Now().Sub(LastMapClear) > MapClearInterval {
			go e.clearOldCachedEvents()
		}

		time.Sleep(time.Second)
	}
}

func (e *eventStreamService) filterEvents(events []data.Event) []data.Event {
	filteredEvents := make([]data.Event, 0)
	for _, event := range events {
		if _, ok := e.fetchedDbEvents[event.ID]; !ok {
			filteredEvents = append(filteredEvents, event)
		}
	}

	return filteredEvents
}

func (e *eventStreamService) clearOldCachedEvents() {
	LastMapClear = time.Now()
	for k, v := range e.fetchedDbEvents {
		if time.Now().Sub(v) > time.Second*30 {
			delete(e.fetchedDbEvents, k)
		}
	}
}

func (e *eventStreamService) clearOldSessions() {
	LastSessionClear = time.Now()

	// End stale sessions
	sessions, err := e.repo.GetStaleSessions()
	if err != nil {
		return
	}
	for _, sessionID := range *sessions {
		err = e.repo.EndSession(sessionID)
		if err != nil {
			continue
		}
	}

	// Ensure ended sessions are not in the stream
	for _, sessionID := range *sessions {
		// Notify streams of session end and remove from stream map
		if _, ok := e.streamMap[sessionID]; ok {
			var conns []*websocket.Conn
			for _, stream := range e.streamMap[sessionID] {
				_ = stream.conn.WriteControl(websocket.CloseMessage, []byte("  session ended"), time.Now().Add(time.Second))
				conns = append(conns, stream.conn)
			}
			delete(e.streamMap, sessionID)

			// Sleep one second to allow the close message to be sent
			time.Sleep(time.Second)
			for _, conn := range conns {
				_ = conn.Close()
			}
		}
	}
}

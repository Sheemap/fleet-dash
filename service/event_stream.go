package service

import (
	"fleet-dash-core/data"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/websocket"
	"time"
)

type stream struct {
	sessionID string
	conn      *websocket.Conn
}

type eventStreamService struct {
	logger          log.Logger
	repo            data.Repository
	streams         []stream
	register        chan stream
	events          chan *[]data.Event
	fetchedDbEvents map[string]time.Time
}

type EventStreamService interface {
	GenerateEventStreamTicket(sessionId string) (*string, error)
	GetActiveTicket(ticket string) (*data.EventStreamTicket, error)
	RegisterStream(sessionID string, conn *websocket.Conn) error
}

func NewEventStreamService(r data.Repository, l log.Logger) EventStreamService {
	svc := &eventStreamService{
		logger:          l,
		repo:            r,
		streams:         make([]stream, 0),
		register:        make(chan stream),
		events:          make(chan *[]data.Event, 100),
		fetchedDbEvents: make(map[string]time.Time),
	}

	go svc.startDBPoll()
	go svc.startStreams()

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

func (e *eventStreamService) RegisterStream(sessionID string, conn *websocket.Conn) error {
	e.register <- stream{sessionID, conn}
	return nil
}

func (e *eventStreamService) startStreams() {
	for {
		select {
		case s := <-e.register:
			e.streams = append(e.streams, s)
		case events := <-e.events:
			if len(*events) == 0 || len(e.streams) == 0 {
				continue
			}
			erroredIndexes := make([]int, 0)
			for i, s := range e.streams {
				// Filter all events to the stream's session
				sessionEvents := make([]data.Event, 0)
				for _, event := range *events {
					if event.SessionID == s.sessionID {
						sessionEvents = append(sessionEvents, event)
					}
				}

				// If any events were filtered, send them to the stream
				if len(sessionEvents) > 0 {
					err := s.conn.WriteJSON(sessionEvents)
					if err != nil {
						erroredIndexes = append(erroredIndexes, i)
					}
				}
			}

			// Remove errored streams from the slice
			for j, i := range erroredIndexes {
				// Slice is getting smaller, so we need to adjust the index
				i = i - j
				if i == len(erroredIndexes) {
					e.streams = e.streams[:i]
				} else {
					e.streams = append(e.streams[:i], e.streams[i+1:]...)
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
	// This has to be separate from above loop so a distributed architecture can be used
	endedSessions, err := e.repo.GetRecentEndedSessions()
	for _, sessionID := range *endedSessions {
		for _, stream := range e.streams {
			if stream.sessionID == sessionID {
				err = stream.conn.WriteControl(websocket.CloseMessage, []byte("  session ended"), time.Now().Add(time.Second))
				if err != nil {
					continue
				}

				time.Sleep(time.Second)
				_ = stream.conn.Close()
			}
		}
	}
}

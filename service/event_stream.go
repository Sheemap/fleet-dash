package service

import (
	"fleet-dash-core/data"
	"github.com/gorilla/websocket"
	"time"
)

type stream struct {
	sessionID string
	conn *websocket.Conn
}

type eventStreamService struct{
	repo data.Repository
	streams []stream
	register chan stream
	events chan *data.Event
}

type EventStreamService interface {
	GenerateEventStreamTicket(sessionId string) (*string, error)
	GetActiveTicket(ticket string) (*data.EventStreamTicket, error)
	RegisterStream(sessionID string, conn *websocket.Conn) error
}

func NewEventStreamService(r data.Repository) EventStreamService {
	svc := &eventStreamService{
		repo: r,
		streams: make([]stream, 0),
		register: make(chan stream),
		events: make(chan *data.Event, 100),
	}

	go svc.startDBPoll()
	go svc.startStreams()

	return svc
}

func (e *eventStreamService)GenerateEventStreamTicket(sessionId string) (*string, error){
	ticket, err := e.repo.GenerateEventStreamTicket(sessionId)
	if err != nil {
		return nil, err
	}

	return &ticket.ID, nil
}

func (e *eventStreamService) GetActiveTicket(ticket string) (*data.EventStreamTicket, error){
	activeTicket, err := e.repo.GetActiveTicket(ticket)
	if err != nil {
		return nil, err
	}

	return activeTicket, nil
}

func (e *eventStreamService) RegisterStream(sessionID string, conn *websocket.Conn) error{
	e.register <- stream{sessionID, conn}
	return nil
}

func (e *eventStreamService) startStreams() {
	for {
		select {
		case s := <-e.register:
			e.streams = append(e.streams, s)
		case event := <-e.events:
			erroredIndexes := make([]int, 0)
			for i, s := range e.streams {
				if s.sessionID == event.SessionID {
					err := s.conn.WriteJSON(event)
					if err != nil {
						erroredIndexes = append(erroredIndexes, i)
					}
				}
			}

			// Remove errored streams from the slice
			for j, i := range erroredIndexes {
				// Slice is getting smaller, so we need to adjust the index
				i = i-j
				if i == len(erroredIndexes){
					e.streams = e.streams[:i]
				} else {
					e.streams = append(e.streams[:i], e.streams[i+1:]...)
				}
			}
		}
	}
}

var SessionClearInterval = time.Minute * 5
func (e *eventStreamService) startDBPoll() {
	lastPoll := time.Now()
	lastSessionClear := time.Now()
	for {
		tmpLastPoll := lastPoll
		lastPoll = time.Now()
		events, err := e.repo.GetEvents(tmpLastPoll)
		if err != nil {
			continue
		}

		for _, event := range events {
			e.events <- event
		}

		if time.Now().Sub(lastSessionClear) > SessionClearInterval {
			// Kick it off in a goroutine so we don't block the main loop
			go func() {
				lastSessionClear = time.Now()

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
			}()
		}

		time.Sleep(time.Second)
	}
}
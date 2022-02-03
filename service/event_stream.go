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

func (e *eventStreamService) startDBPoll() {
	lastPoll := time.Now()
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

		time.Sleep(time.Second)
	}
}
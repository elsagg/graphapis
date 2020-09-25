package events

import (
	"context"

	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/protocol"
)

type EventData struct {
	EventID     string `json:"ce_id"`
	EventType   string `json:"ce_type"`
	EventSource string `json:"ce_source"`
}

func (e *EventData) AutoFill(e2 *Event) {
	e.EventID = e2.ID()
	e.EventType = e2.Type()
	e.EventSource = e2.Source()
}

type Event struct {
	cloudevents.Event
	destination string
	eventKey    string
	client      cloudevents.Client
	sender      *kafka_sarama.Sender
	ctx         context.Context
	sendResult  protocol.Result
}

func (e *Event) Destination() string {
	return e.destination
}

func (e *Event) SetDestination(destination string) {
	e.destination = destination
}

func (e *Event) Key() string {
	return e.eventKey
}

func (e *Event) SetKey(key string) {
	e.eventKey = key
}

func (e *Event) SetEventData(data interface{}) error {
	return e.SetData(cloudevents.ApplicationJSON, data)
}

func NewEvent(ctx context.Context, destination string) (*Event, error) {
	e := Event{
		destination: destination,
		ctx:         ctx,
	}

	e.SetSpecVersion(cloudevents.VersionV1)

	err := e.BuildSender()

	if err != nil {
		return nil, err
	}

	return &e, nil
}

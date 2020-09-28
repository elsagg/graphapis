package events

import (
	"context"
	"time"

	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/protocol"
	uuid "github.com/satori/go.uuid"
)

type EventData struct {
	EventID     string    `json:"ce_id"`
	EventType   string    `json:"ce_type"`
	EventSource string    `json:"ce_source"`
	EventTime   time.Time `json:"ce_time"`
	Event       *Event    `json:"-"`
}

func (e *EventData) CreateEvent(ctx context.Context, metadata *EventMetadata) error {
	ev, err := NewEvent(ctx, metadata)

	if err != nil {
		return err
	}

	e.Event = ev

	e.EventID = ev.ID()
	e.EventType = ev.Type()
	e.EventTime = ev.Time()
	e.EventSource = ev.Source()

	return nil
}

func (e *EventData) SetEvent(ev *Event) {
	e.Event = ev

	e.EventID = ev.ID()
	e.EventType = ev.Type()
	e.EventTime = ev.Time()
	e.EventSource = ev.Source()
}

func (e *EventData) AutoFill(e2 *Event) {
	e.EventID = e2.ID()
	e.EventType = e2.Type()
	e.EventSource = e2.Source()
}

type EventMetadata struct {
	EventType        string
	EventSource      string
	EventKey         string
	EventDestination string
	EventTime        time.Time
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

func NewEvent(ctx context.Context, metadata *EventMetadata) (*Event, error) {
	eventId := uuid.NewV4()

	e := Event{
		ctx: ctx,
	}

	e.SetID(eventId.String())
	e.SetType(metadata.EventType)
	e.SetSource(metadata.EventSource)
	e.SetTime(metadata.EventTime)
	e.SetKey(metadata.EventKey)
	e.SetDestination(metadata.EventDestination)

	e.SetSpecVersion(cloudevents.VersionV1)

	err := e.BuildSender()

	if err != nil {
		return nil, err
	}

	return &e, nil
}

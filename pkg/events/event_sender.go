package events

import (
	"errors"
	"os"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func (e *Event) BuildSender() error {
	brokerList := os.Getenv("KAFKA_BROKERS")
	brokers := strings.Split(brokerList, ",")

	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0
	sender, err := kafka_sarama.NewSender(brokers, saramaConfig, e.Destination())
	if err != nil {
		return err
	}

	e.sender = sender

	client, err := cloudevents.NewClient(sender, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		return err
	}

	e.client = client

	return nil
}

func (e *Event) Send() error {
	result := e.client.Send(
		kafka_sarama.WithMessageKey(e.ctx, sarama.StringEncoder(e.Key())),
		e.Clone(),
	)

	e.sendResult = result

	if cloudevents.IsUndelivered(e.sendResult) {
		err := errors.New("Failed to deliver event")
		return err
	}
	KafkaEventDelivered(e)
	return nil
}

func (e *Event) IsACK() bool {
	return cloudevents.IsACK(e.sendResult)
}

func (e *Event) Close() {
	defer e.sender.Close(e.ctx)
}

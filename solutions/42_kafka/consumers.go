// GoForGo Solution: Kafka Consumers
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(brokers []string, topic, groupID string) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	return &KafkaConsumer{reader: reader}
}

func (kc *KafkaConsumer) consumeMessages(ctx context.Context, messageHandler func(kafka.Message) error) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			msg, err := kc.reader.ReadMessage(ctx)
			if err != nil {
				return err
			}
			if err := messageHandler(msg); err != nil {
				log.Printf("Error handling message: %v", err)
			}
		}
	}
}

func (kc *KafkaConsumer) close() error {
	return kc.reader.Close()
}

func main() {
	brokers := []string{"localhost:9092"}
	consumer := NewKafkaConsumer(brokers, "user-events", "test-group")
	defer consumer.close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	consumer.consumeMessages(ctx, func(msg kafka.Message) error {
		log.Printf("Consumed: %s = %s", string(msg.Key), string(msg.Value))
		return nil
	})

	fmt.Println("Kafka consumer operations completed!")
}
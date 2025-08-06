// GoForGo Solution: Kafka Streams
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Event struct {
	UserID    string    `json:"user_id"`
	EventType string    `json:"event_type"`
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
}

type StreamProcessor struct {
	reader *kafka.Reader
	writer *kafka.Writer
}

func NewStreamProcessor(inputTopic, outputTopic string, brokers []string) *StreamProcessor {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   inputTopic,
		GroupID: "stream-processor",
	})
	
	writer := &kafka.Writer{
		Addr:  kafka.TCP(brokers...),
		Topic: outputTopic,
	}

	return &StreamProcessor{reader: reader, writer: writer}
}

func (sp *StreamProcessor) processStream(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			msg, err := sp.reader.ReadMessage(ctx)
			if err != nil {
				return err
			}

			var event Event
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				log.Printf("Error unmarshaling event: %v", err)
				continue
			}

			// Filter and transform
			if event.EventType == "purchase" {
				processedEvent := map[string]interface{}{
					"user_id":     event.UserID,
					"processed":   true,
					"timestamp":   time.Now(),
					"original":    event,
				}

				processedData, _ := json.Marshal(processedEvent)
				sp.writer.WriteMessages(ctx, kafka.Message{
					Key:   msg.Key,
					Value: processedData,
				})
				
				log.Printf("Processed purchase event for user: %s", event.UserID)
			}
		}
	}
}

func main() {
	processor := NewStreamProcessor("input-events", "processed-events", []string{"localhost:9092"})
	
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	processor.processStream(ctx)
	fmt.Println("Kafka streams operations completed!")
}
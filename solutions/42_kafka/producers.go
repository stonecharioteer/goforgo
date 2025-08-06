// GoForGo Solution: Kafka Producers
// Complete implementation of Kafka message production with various publishing patterns

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// KafkaProducer handles message publishing to Kafka topics
type KafkaProducer struct {
	writer *kafka.Writer
	topic  string
}

// MessagePair represents a key-value message pair
type MessagePair struct {
	Key   string
	Value string
}

// NewKafkaProducer creates a new Kafka producer
func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{}, // Distribute messages evenly
		BatchTimeout: 10 * time.Millisecond,
		BatchSize:    100,
	}

	log.Printf("Created Kafka producer for topic: %s", topic)
	return &KafkaProducer{
		writer: writer,
		topic:  topic,
	}
}

// sendMessage publishes a single message to the topic
func (kp *KafkaProducer) sendMessage(ctx context.Context, key, value string) error {
	message := kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
		Time:  time.Now(),
	}

	err := kp.writer.WriteMessages(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	log.Printf("Sent message - Key: %s, Value: %s", key, value)
	return nil
}

// sendBatchMessages publishes multiple messages efficiently
func (kp *KafkaProducer) sendBatchMessages(ctx context.Context, messages []MessagePair) error {
	kafkaMessages := make([]kafka.Message, len(messages))
	
	for i, msg := range messages {
		kafkaMessages[i] = kafka.Message{
			Key:   []byte(msg.Key),
			Value: []byte(msg.Value),
			Time:  time.Now(),
		}
	}

	err := kp.writer.WriteMessages(ctx, kafkaMessages...)
	if err != nil {
		return fmt.Errorf("failed to write batch messages: %w", err)
	}

	log.Printf("Sent batch of %d messages", len(messages))
	return nil
}

// sendMessageWithHeaders publishes a message with custom headers
func (kp *KafkaProducer) sendMessageWithHeaders(ctx context.Context, key, value string, headers map[string]string) error {
	kafkaHeaders := make([]kafka.Header, 0, len(headers))
	for k, v := range headers {
		kafkaHeaders = append(kafkaHeaders, kafka.Header{
			Key:   k,
			Value: []byte(v),
		})
	}

	message := kafka.Message{
		Key:     []byte(key),
		Value:   []byte(value),
		Headers: kafkaHeaders,
		Time:    time.Now(),
	}

	err := kp.writer.WriteMessages(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to write message with headers: %w", err)
	}

	log.Printf("Sent message with headers - Key: %s, Value: %s, Headers: %v", key, value, headers)
	return nil
}

// sendMessageToPartition publishes a message to a specific partition
func (kp *KafkaProducer) sendMessageToPartition(ctx context.Context, partition int, key, value string) error {
	message := kafka.Message{
		Key:       []byte(key),
		Value:     []byte(value),
		Partition: partition,
		Time:      time.Now(),
	}

	err := kp.writer.WriteMessages(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to write message to partition %d: %w", partition, err)
	}

	log.Printf("Sent message to partition %d - Key: %s, Value: %s", partition, key, value)
	return nil
}

// sendMessageWithTimestamp publishes a message with a specific timestamp
func (kp *KafkaProducer) sendMessageWithTimestamp(ctx context.Context, key, value string, timestamp time.Time) error {
	message := kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
		Time:  timestamp,
	}

	err := kp.writer.WriteMessages(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to write message with timestamp: %w", err)
	}

	log.Printf("Sent message with timestamp %s - Key: %s, Value: %s", timestamp.Format(time.RFC3339), key, value)
	return nil
}

// close closes the Kafka writer and releases resources
func (kp *KafkaProducer) close() error {
	if err := kp.writer.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}
	log.Println("Kafka producer closed successfully")
	return nil
}

// generateSampleData creates sample messages for testing
func generateSampleData(count int) []MessagePair {
	messages := make([]MessagePair, count)
	
	events := []string{"login", "logout", "purchase", "view_product", "add_to_cart"}
	
	for i := 0; i < count; i++ {
		messages[i] = MessagePair{
			Key:   fmt.Sprintf("user_%d", i%100),
			Value: fmt.Sprintf(`{"event": "%s", "user_id": "user_%d", "timestamp": "%s"}`, 
				events[i%len(events)], i%100, time.Now().Format(time.RFC3339)),
		}
	}
	
	log.Printf("Generated %d sample messages", count)
	return messages
}

func main() {
	// Define Kafka broker addresses
	brokers := []string{"localhost:9092"}
	topic := "user-events"

	// Create a Kafka producer
	producer := NewKafkaProducer(brokers, topic)
	defer producer.close()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Send a single message
	if err := producer.sendMessage(ctx, "user123", "user logged in"); err != nil {
		log.Printf("Error sending single message: %v", err)
	}

	// Send a message with headers
	headers := map[string]string{
		"event-type": "login",
		"source":     "web-app",
		"version":    "1.0",
	}
	if err := producer.sendMessageWithHeaders(ctx, "user456", "authenticated user login", headers); err != nil {
		log.Printf("Error sending message with headers: %v", err)
	}

	// Generate and send batch messages
	sampleMessages := generateSampleData(10)
	if err := producer.sendBatchMessages(ctx, sampleMessages); err != nil {
		log.Printf("Error sending batch messages: %v", err)
	}

	// Send a message to a specific partition (partition 0)
	if err := producer.sendMessageToPartition(ctx, 0, "admin-action", "user created new account"); err != nil {
		log.Printf("Error sending message to partition: %v", err)
	}

	// Send a message with custom timestamp (1 hour ago)
	historicalTime := time.Now().Add(-1 * time.Hour)
	if err := producer.sendMessageWithTimestamp(ctx, "historical-event", "system maintenance completed", historicalTime); err != nil {
		log.Printf("Error sending message with timestamp: %v", err)
	}

	// Demonstrate error handling - try to send to a topic that requires different configuration
	log.Println("Demonstrating error handling...")
	errorCtx, errorCancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer errorCancel()
	
	if err := producer.sendMessage(errorCtx, "test", "this might timeout"); err != nil {
		log.Printf("Expected error (timeout or connection): %v", err)
	}

	// Wait a moment for all messages to be sent
	time.Sleep(2 * time.Second)

	fmt.Println("Kafka producer operations completed!")
}
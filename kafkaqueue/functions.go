package kafkaqueue

import (
	"context"
	"log"
	"strconv"

	"github.com/segmentio/kafka-go"
)

const (
    kafkaBroker = "localhost:9092" 
    topic       = "test" 
)

// Producer function to publish the number to Kafka
func PublishToKafka(number int) {
    w := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{kafkaBroker},
        Topic:   topic,
    })
    defer w.Close()

    message := strconv.Itoa(number)
    err := w.WriteMessages(context.Background(), kafka.Message{
        Value: []byte(message),
    })

    if err != nil {
        log.Println("Failed to send message:", err)
    } else {
        log.Println("Message sent successfully")
    }
}

// Consumer function to read the value from Kafka
func ConsumeFromKafka() int {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{kafkaBroker},
        Topic:   topic,
    })
    defer r.Close()
	var number int
    for {
        msg, err := r.ReadMessage(context.Background())
        if err != nil {
            log.Println("Error while reading message:", err)
            return -1
        }

        value := string(msg.Value)
        number, err = strconv.Atoi(value)
        if err != nil {
            log.Println("Error parsing Kafka message:", err)
            continue
        }
		return number
    }

}

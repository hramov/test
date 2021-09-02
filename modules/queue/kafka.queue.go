package queue

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var (
	Producer *kafka.Producer
	Consumer *kafka.Consumer
	Messages chan string = make(chan string)
)

type Queue struct {
	Producer *kafka.Producer
	Consumer *kafka.Consumer
	Messages chan string
}

func CreateConsumer(topic string) *kafka.Consumer {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "192.168.0.103",
		"group.id":          "test",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Println(err)
		return nil
	}
	defer c.Close()
	c.SubscribeTopics([]string{topic}, nil)
	log.Println("Kafka consumer created")
	Consumer = c
	return c
}

func CreateProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "192.168.0.103"})
	if err != nil {
		log.Println(err)
		return nil
	}
	defer p.Close()
	log.Println("Kafka producer created")
	Producer = p
	return p
}

func (q *Queue) Monitoring() {
	for e := range Producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
			} else {
				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
			}
		}
	}
}

func (q *Queue) SendMessage(topic string, message string) {
	for _, word := range strings.Split(message, " ") {
		Producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
	Producer.Flush(15 * 1000)
}

func (q *Queue) ReceiveMessage() {
	log.Println("Waiting for messages")
	for {
		msg, err := Consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			Messages <- string(msg.Value)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

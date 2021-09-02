package queue

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Queue struct {
	Producer *kafka.Producer
	Consumer *kafka.Consumer
	Messages chan string
}

func CreateConsumer(topic string) *kafka.Consumer {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "test",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Println(err)
		return nil
	}
	defer c.Close()
	c.SubscribeTopics([]string{topic}, nil)
	return c
}

func CreateProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		log.Println(err)
		return nil
	}
	defer p.Close()
	return p
}

func (q *Queue) Monitoring() {
	for e := range q.Producer.Events() {
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
		q.Producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
	q.Producer.Flush(15 * 1000)
}

func (q *Queue) ReceiveMessage() {
	for {
		msg, err := q.Consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			q.Messages <- string(msg.Value)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

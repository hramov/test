package main

import (
	"github.com/golobby/container/v3"
	"log"
	"testing/domains/entities"
	"testing/domains/gates"
	"testing/domains/ports"
	"testing/modules/cache"
	"testing/modules/database"
	"testing/modules/database/access"
	"testing/modules/grpc"
	"testing/modules/logger"
	"testing/modules/queue"
)

func IoCInit() error {
	var postgres gates.Orm = &database.Gorm{}
	postgres.Connect()
	postgres.Migrate()
	pgConnection := postgres.GetConnection()
	chConnection := logger.Connect()
	redisConnection := cache.Connect()

	kafkaConsumer := queue.CreateConsumer("logs")
	kafkaProducer := queue.CreateProducer()
	kafkaChan := make(chan string)
	kafka := &queue.Queue{
		Producer: kafkaProducer,
		Consumer: kafkaConsumer,
		Messages: kafkaChan,
	}
	kafka.SendMessage("logs", "Kafka has been started")

	err := container.NamedSingleton("UserEntity", func() ports.UserEntityPort {
		return &entity.UserEntity{
			Provider: &access.UserProvider{
				DB: pgConnection,
			}}
	})
	err = container.NamedSingleton("Logger", func() ports.LoggerPort {
		return &logger.Logger{
			DB: chConnection,
		}
	})
	err = container.NamedSingleton("Cache", func() ports.CachePort {
		return &cache.Cache{
			Client: redisConnection,
		}
	})

	err = container.NamedSingleton("Queue", func() ports.QueuePort {
		return kafka
	})

	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := IoCInit(); err != nil {
		log.Fatal("Cannot use IoC container!")
	}
	var queue queue.Queue
	container.NamedResolve(&queue, "Queue")

	var logger logger.Logger
	container.NamedResolve(&logger, "Logger")
	logger.Log("Logger has been started")

	// go queue.ReceiveMessage()
	// go logger.Log(<-queue.Messages)

	grpc := grpc.Server{}
	grpc.Start()
}

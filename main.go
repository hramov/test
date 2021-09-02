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
	err := container.NamedSingleton("UserEntity", func() ports.UserEntityPort {
		var postgres gates.Orm = &database.Gorm{}
		postgres.Connect()
		postgres.Migrate()

		return &entity.UserEntity{
			Provider: &access.UserProvider{
				DB: postgres.GetConnection(),
			}}
	})
	err = container.NamedSingleton("Logger", func() ports.LoggerPort {
		return &logger.Logger{
			DB: logger.Connect(),
		}
	})
	err = container.NamedSingleton("Cache", func() ports.CachePort {
		return &cache.Cache{
			Client: cache.Connect(),
		}
	})
	err = container.NamedSingleton("Queue", func() ports.QueuePort {
		return &queue.Queue{
			Producer: queue.CreateProducer(),
			Consumer: queue.CreateConsumer("logs"),
			Messages: make(chan string),
		}
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

	go queue.ReceiveMessage()
	go logger.Log(<-queue.Messages)

	grpc := grpc.Server{}
	grpc.Start()
}

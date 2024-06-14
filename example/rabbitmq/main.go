package main

import (
	"log"
	"time"

	"github.com/jamesstocktonj1/chlogger"
	"github.com/jamesstocktonj1/chlogger/rabbitmq"
)

var logger chlogger.Logger

func main() {
	logger = rabbitmq.NewLogger(rabbitmq.RabbitMQConfig{
		Host:  "amqp://user:password@localhost:5672/",
		Topic: "logger",
		AppID: "example",
	})

	err := logger.Init()
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	for {
		logger.Printf("Hello World! %d", i)
		time.Sleep(time.Second)
		i++
	}
}

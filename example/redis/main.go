package main

import (
	"log"
	"time"

	"github.com/jamesstocktonj1/chlogger"
	"github.com/jamesstocktonj1/chlogger/redis"
)

var logger chlogger.Logger

func main() {
	logger = redis.NewLogger(redis.RedisConfig{
		Host:  "localhost:6379",
		Topic: "logger",
	})
	defer logger.Close()

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

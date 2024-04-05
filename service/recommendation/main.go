package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"literank.com/event-books/infrastructure/parser"
	"literank.com/event-books/service/recommendation/application"
	"literank.com/event-books/service/recommendation/application/consumer"
	"literank.com/event-books/service/recommendation/infrastructure/config"
)

const configFileName = "config.yml"

func main() {
	// Read the config
	c, err := parser.Parse[config.Config](configFileName)
	if err != nil {
		panic(err)
	}

	// Prepare dependencies
	wireHelper, err := application.NewWireHelper(c)
	if err != nil {
		panic(err)
	}

	// Run the consumer
	tc := consumer.NewInterestConsumer(wireHelper.InterestManager(), wireHelper.TrendEventConsumer())
	eventConsumer := tc.EventConsumer()
	go func() {
		// Shutdown signals
		stopAll := make(chan os.Signal, 1)
		signal.Notify(stopAll, syscall.SIGINT, syscall.SIGTERM)
		<-stopAll
		if err := eventConsumer.Stop(); err != nil {
			log.Panicf("Failed to close consumer group: %v", err)
		}
	}()
	tc.Start(context.Background())
}

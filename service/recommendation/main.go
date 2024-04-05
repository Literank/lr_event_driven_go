package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"literank.com/event-books/infrastructure/parser"
	"literank.com/event-books/service/recommendation/adapter"
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
		tc.Start(context.Background())

	}()

	// Build main router
	r, err := adapter.MakeRouter(wireHelper)
	if err != nil {
		panic(err)
	}

	svr := &http.Server{
		Addr:    fmt.Sprintf(":%d", c.App.Port),
		Handler: r,
	}

	// Shutdown signals
	stopAll := make(chan os.Signal, 1)
	signal.Notify(stopAll, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-stopAll
		if err := eventConsumer.Stop(); err != nil {
			log.Panicf("Failed to close consumer group: %v", err)
		}
		if err := svr.Shutdown(context.Background()); err != nil {
			log.Panicf("Failed to shutdown Gin server: %v", err)
		}
	}()

	// Run the server on the specified port
	if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

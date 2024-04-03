package main

import (
	"fmt"

	"literank.com/event-books/infrastructure/parser"
	"literank.com/event-books/service/trend/adapter"
	"literank.com/event-books/service/trend/application"
	"literank.com/event-books/service/trend/infrastructure/config"
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

	// Build main router
	r, err := adapter.MakeRouter(wireHelper)
	if err != nil {
		panic(err)
	}
	// Run the server on the specified port
	if err := r.Run(fmt.Sprintf(":%d", c.App.Port)); err != nil {
		panic(err)
	}
}

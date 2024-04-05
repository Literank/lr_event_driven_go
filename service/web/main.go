package main

import (
	"fmt"

	"literank.com/event-books/infrastructure/parser"
	"literank.com/event-books/service/web/adapter"
	"literank.com/event-books/service/web/application"
	"literank.com/event-books/service/web/infrastructure/config"
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
	r, err := adapter.MakeRouter(c.App.TemplatesPattern, &c.Remote, wireHelper)
	if err != nil {
		panic(err)
	}
	// Run the server on the specified port
	if err := r.Run(fmt.Sprintf(":%d", c.App.Port)); err != nil {
		panic(err)
	}
}

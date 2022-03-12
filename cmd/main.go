package main

import (
	"context"
	"log"

	"history-rate/app"
)

func main() {
	ctx := context.Background()
	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

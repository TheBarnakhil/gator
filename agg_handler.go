package main

import (
	"context"
	"fmt"
	"log"
)

func handleAgg(_ *state, _ command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		log.Fatalf("error fetching the feed: %v", err)
	}
	fmt.Println(feed)
	return nil
}

package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/TheBarnakhil/gator/internal/database"
	"github.com/google/uuid"
)

func handleAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("login expects only 1 argument, time_between_reqs")
	}
	duration, err := time.ParseDuration(string(cmd.args[0]))
	if err != nil {
		return fmt.Errorf("error parsing the duration: %v", err)
	}
	println("Collecting feeds every", duration)
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
	}
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("error marking the feed: %v", err)
		return
	}
	fetched, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("error fetching the feed: %v", err)
		return
	}
	for _, item := range fetched.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			FeedID:    feed.ID,
			Title:     item.Title,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			Url:         item.Link,
			PublishedAt: publishedAt,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(fetched.Channel.Item))
}

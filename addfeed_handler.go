package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/TheBarnakhil/gator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return errors.New("addfeed expects two arguments: name and url")
	}
	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      cmd.args[0],
			Url:       cmd.args[1],
			UserID:    user.ID,
		},
	)
	if err != nil {
		log.Fatalf("Error creating feed: %v", err)
	}
	fmt.Printf("Feed %v has been created!\n", feed)
	follow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}
	fmt.Printf("User %s is now following the feed %s !\n", follow.UserName, follow.FeedName)
	return nil
}

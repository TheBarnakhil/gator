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

func handleFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("follow expects only 1 argument, url of feed")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		log.Fatalf("Error feed doesn't exist!: %v", cmd.args[0])
	}
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

package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/TheBarnakhil/gator/internal/database"
)

func handleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("follow expects only 1 argument, url of feed")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		log.Fatalf("Error feed doesn't exist!: %v", cmd.args[0])
	}
	err = s.db.DeleteFollow(
		context.Background(),
		database.DeleteFollowParams{
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}
	fmt.Printf("User %s has now unfollowed the feed %s !\n", user.Name, feed.Name)
	return nil
}

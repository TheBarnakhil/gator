package main

import (
	"context"
	"fmt"

	"github.com/TheBarnakhil/gator/internal/database"
)

func handleFollowing(s *state, _ command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting follows for current user: %w", err)
	}
	for _, follow := range follows {
		fmt.Println(follow.FeedName)
	}
	return nil
}

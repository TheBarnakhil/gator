package main

import (
	"context"
	"fmt"
)

func handleFollowing(s *state, _ command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting current user info: %w", err)
	}
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting follows for current user: %w", err)
	}
	for _, follow := range follows {
		fmt.Println(follow.FeedName)
	}
	return nil
}

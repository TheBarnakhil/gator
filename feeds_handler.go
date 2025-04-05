package main

import (
	"context"
	"fmt"
	"log"
)

func handleFeeds(s *state, _ command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	for i, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error %w retrieving user for feed %s", err, feed.Name)
		}
		fmt.Printf("Feed %d :\n", i+1)
		fmt.Println("  - Name :", feed.Name)
		fmt.Println("  - URL :", feed.Name)
		fmt.Println("  - Created By :", user.Name.String)
	}
	if err != nil {
		log.Fatalf("Error deleting all users!")
	}
	fmt.Println("All user records have been deleted!")
	return nil
}

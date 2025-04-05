package main

import (
	"context"
	"fmt"
	"log"
)

func handleUsers(s *state, _ command) error {
	users, err := s.db.GetUsers(context.Background())
	for _, user := range users {
		fmt.Print(user.Name)
		if user.Name == s.cfg.CurrentUserName {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}
	if err != nil {
		log.Fatalf("Error deleting all users!")
	}
	fmt.Println("All user records have been deleted!")
	return nil
}

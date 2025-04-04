package main

import (
	"context"
	"fmt"
	"log"
)

func handleReset(s *state, _ command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		log.Fatalf("Error deleting all users!")
	}
	fmt.Println("All user records have been deleted!")
	return nil
}

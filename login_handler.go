package main

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("login expects only 1 argument, username")
	}
	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		log.Fatalf("Error user doesn't exist!: %v", cmd.args[0])
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		log.Fatalf("Error setting user: %v", err)
	}
	fmt.Printf("User %s has been set!\n", user.Name)
	return nil
}

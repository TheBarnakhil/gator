package main

import (
	"errors"
	"fmt"
	"log"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("login expects only 1 argument, username")
	}
	err := s.cfg.SetUser(string(cmd.args[0]))
	if err != nil {
		log.Fatalf("Error setting user: %v", err)
	}
	fmt.Printf("User %s has been set!\n", cmd.args[0])
	return nil
}

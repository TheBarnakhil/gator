package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("login expects only 1 argument, username")
	}
	user, err := s.db.GetUser(context.Background(), sql.NullString{
		String: cmd.args[0],
		Valid:  true,
	})
	if err != nil {
		log.Fatalf("Error user doesn't exist!: %v", cmd.args[0])
	}
	err = s.cfg.SetUser(user.Name.String)
	if err != nil {
		log.Fatalf("Error setting user: %v", err)
	}
	fmt.Printf("User %s has been set!\n", user.Name.String)
	return nil
}

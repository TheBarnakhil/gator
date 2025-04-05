package main

import (
	"context"
	"fmt"

	"github.com/TheBarnakhil/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error user doesn't exist!: %v", s.cfg.CurrentUserName)
		}
		return handler(s, c, user)
	}
}

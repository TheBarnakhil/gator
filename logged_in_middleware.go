package main

import "github.com/TheBarnakhil/gator/internal/database"

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error { return nil }
}

package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/TheBarnakhil/gator/internal/database"
	"github.com/google/uuid"
)

func handleRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("register expects only 1 argument, username")
	}
	user, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: sql.NullString{
				String: cmd.args[0],
				Valid:  true,
			},
		},
	)
	if err != nil {
		log.Fatalf("Error setting user: %v", err)
	}
	err = s.cfg.SetUser(user.Name.String)
	if err != nil {
		log.Fatalf("Error setting user: %v", err)
	}
	fmt.Printf("User %v has been set!\n", user)
	return nil
}

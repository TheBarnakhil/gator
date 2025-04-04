package main

import (
	"github.com/TheBarnakhil/gator/internal/config"
	"github.com/TheBarnakhil/gator/internal/database"
)

type (
	state struct {
		cfg *config.Config
		db  *database.Queries
	}

	command struct {
		name string
		args []string
	}
)

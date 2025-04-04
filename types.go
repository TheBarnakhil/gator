package main

import "github.com/TheBarnakhil/gator/internal/config"

type (
	state struct {
		cfg *config.Config
	}

	command struct {
		name string
		args []string
	}
)

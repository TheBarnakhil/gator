package main

import (
	"fmt"
)

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, exists := c.cmds[cmd.name]
	if !exists {
		return fmt.Errorf("Invalid command: %s", cmd.name)
	}
	return f(s, cmd)
}

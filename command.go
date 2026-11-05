package main

import "fmt"

// What the user typed: e.g. login alice → Name "login", Args ["alice"]
type command struct {
	Name string
	Args []string
}

// Maps a command name to the function that handles it
type commands struct {
	registeredCommands map[string]func(*state, command) error
}

// Stores the handler under its name (no () after f, so it isn't called here)
func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

// Looks up the handler by name and calls it
func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return f(s, cmd)
}

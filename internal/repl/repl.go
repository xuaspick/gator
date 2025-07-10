package repl

import (
	"fmt"

	"github.com/xuaspick/gator/internal/config"
)

type state struct {
	Config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	m map[string]func(*state, command) error
}

func GetCommands() *commands {
	return &commands{
		m: map[string]func(*state, command) error{
			"login": HandlerLogin,
		},
	}
}

func HandlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Login command expects the username argument")
	}

	err := s.Config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("%v has been set as userName", s.Config.CurrentUserName)
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.m[cmd.name]
	if !ok {
		return fmt.Errorf("The comand %v is not registered", cmd.name)
	}
	return f(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.m[name] = f
}

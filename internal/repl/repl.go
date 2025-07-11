package repl

import (
	"fmt"

	"github.com/xuaspick/gator/internal/config"
	"github.com/xuaspick/gator/internal/database"
)

type State struct {
	DB  *database.Queries
	Cfg *config.Config
}

type Command struct {
	Name string
	Args []string
}

type commands struct {
	m map[string]func(*State, Command) error
}

func GetCommands() *commands {
	return &commands{
		m: map[string]func(*State, Command) error{},
	}
}

// func HandlerLogin(s *State, cmd Command) error {
// 	if len(cmd.Args) == 0 {
// 		return fmt.Errorf("Login command expects the username argument")
// 	}

// 	err := s.Cfg.SetUser(cmd.Args[0])
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("%v has been set as userName\n", s.Cfg.CurrentUserName)
// 	return nil
// }

func (c *commands) Run(s *State, cmd Command) error {
	f, ok := c.m[cmd.Name]
	if !ok {
		return fmt.Errorf("The comand %v is not registered", cmd.Name)
	}
	return f(s, cmd)
}

func (c *commands) Register(name string, f func(*State, Command) error) {
	c.m[name] = f
}

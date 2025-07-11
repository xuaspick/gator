package repl

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/xuaspick/gator/internal/database"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	user, err := s.DB.GetUser(context.Background(), cmd.Args[0])

	if err != nil {
		return err
	}

	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("%v has been set as userName\n", s.Cfg.CurrentUserName)
	return nil
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		return err
	}

	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("user registered \n %v \n", user)

	return nil
}

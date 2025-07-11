package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/xuaspick/gator/internal/config"
	"github.com/xuaspick/gator/internal/database"
	"github.com/xuaspick/gator/internal/repl"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}
	state := &repl.State{
		Cfg: &cfg,
	}
	cmds := repl.GetCommands()
	cmds.Register("login", repl.HandlerLogin)
	cmds.Register("register", repl.HandlerRegister)

	if len(os.Args) < 2 {
		log.Fatal("CLI expects at least 1 argument to be passed")
	}

	cmd := repl.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}
	db, err := sql.Open("postgres", state.Cfg.DBURL)
	state.DB = database.New(db)

	err = cmds.Run(state, cmd)
	if err != nil {
		log.Fatal(err)
	}

}

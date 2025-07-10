package main

import (
	"fmt"
	"log"

	"github.com/xuaspick/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}
	// cfg.SetUser("xuaspick")

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	fmt.Printf("%+v \n", cfg)

}

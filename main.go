package main

import (
	"log"

	core "github.com/deifyed/post-service/pkg/core"

	"fmt"
	"os"
)

func main() {
	log.Printf("Loading config")

	cfg, err := core.LoadConfig(os.Getenv)
	if err != nil {
		log.Fatal(fmt.Errorf("loading config: %w", err))
	}

	log.Printf("Loaded config: %+v", cfg)

	log.Printf("Server started")

	router := core.NewRouter(cfg)

	log.Fatal(router.Run(fmt.Sprintf(":%d", cfg.Port)))
}

package main

import (
	"log"

	sw "github.com/deifyed/post-service/pkg/core"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(router.Run(":3000"))
}

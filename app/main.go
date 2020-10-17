package main

import (
	"log"

	"github.com/hokita/routine/http"
)

func main() {
	if err := http.Start(); err != nil {
		log.Fatalf("could not listen on port 8081 %v", err)
	}
}

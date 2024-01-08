package main

import (
	"github.com/Alexander2k/CryptoBotGo/internal/metrics"
	"log"
)

func main() {

	go func() {
		err := metrics.Listen("localhost:8989")
		if err != nil {
			log.Printf("Error listening prometheus metrics: %v", err)
		}
	}()

	if err := Start(); err != nil {
		log.Fatal(err)
	}
}

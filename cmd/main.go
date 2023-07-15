package main

import "log"

func main() {
	if err := Start(); err != nil {
		log.Fatal(err)
	}
}

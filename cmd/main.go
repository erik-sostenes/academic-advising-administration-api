package main

import (
	"log"

	"github.com/itsoeh/academy-advising-administration-api/internal/dependency"
)

func main() {
	// start running the program
	if err := dependency.Run(); err != nil {
		log.Fatal(err)
	}
} 

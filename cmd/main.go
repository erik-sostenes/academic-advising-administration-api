package main

import (
	"log"

	"github.com/itsoeh/academic-advising-administration-api/internal/dependency"
)

func main() {
	// start running the program
	if err := dependency.Run(); err != nil {
		log.Fatal(err)
	}
} 

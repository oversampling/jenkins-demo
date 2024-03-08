package main

// Import the fmt for formatting strings
// Import os so we can read environment variables from the system
import (
	"log"
	"time"
)

func main() {
	log.Println("Welcome to use jenkins test test test")
	ticker := time.NewTicker(60 * time.Second)
	for range ticker.C {
		log.Println("Hello, Kubernetes！I'm from Jenkins CI！")
		log.Println("Hello everyone!!!!")
	}
}

package main

// Import the fmt for formatting strings
// Import os so we can read environment variables from the system
import (
	"log"
	"time"
)

func main() {
	log.Println("Welcome for Jenkins POC Demo 100")
	ticker := time.NewTicker(60 * time.Second)
	for range ticker.C {
		log.Println("Hello, CCE Kubernetes Jenkins CI！")
	}
}

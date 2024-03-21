package main

// Import the fmt for formatting strings
// Import os so we can read environment variables from the system
import (
	"log"
	"time"
)

func main() {
	log.Println("Hi LikeCard, Huawei Cloud Definitely Can Run CICD pipeline like other cloud provider")
	ticker := time.NewTicker(60 * time.Second)
	for range ticker.C {
		log.Println("Likecard")
	}
}

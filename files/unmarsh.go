package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string
	Description string
	Age         int
}

func main() {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks", "Age":29}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s, Age: %d", bird.Species, bird.Description, bird.Age)

}

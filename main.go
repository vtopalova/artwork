package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Write a function that concurrently fetches a list of artworks (given by their IDs) and returns their titles as a slice of strings.
// Expected result:
// Fetching artwork 101...
// Fetching artwork 202...
// Fetching artwork 303...
// Fetching artwork 404...
// Done: [Artwork Title 101 Artwork Title 202 Artwork Title 303 Artwork Title 404]

// Simulated API call (we pretend to fetch artwork title by ID)
func fetchArtwork(id int) string {
	fmt.Printf("Fetching artwork %d...\n", id)
	time.Sleep(time.Duration(rand.Intn(900)+100) * time.Millisecond)
	return fmt.Sprintf("Artwork Title %d", id)
}

// TODO: Implement this function
func fetchArtworks(ids []int) []string {
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ids := []int{101, 202, 303, 404}
	titles := fetchArtworks(ids)
	fmt.Println("Done:", titles)
}

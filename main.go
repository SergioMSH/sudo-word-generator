package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate a random word with 10 characters
	randomWord := generateRandomWord(10)
	// Print the random word to the console
	fmt.Println(randomWord)
}

func generateRandomWord(length int) string {
	// Define a slice of runes containing all the letters of the alphabet
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())
	// Create a slice of runes with the specified length
	b := make([]rune, length)
	// Fill the slice with random letters from the letters slice
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	// Convert the slice of runes to a string and return it
	return string(b)
}

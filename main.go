package main

import (
	"fmt"
	"os"

	"main.go/git"
)

func main() {
	// Get git diff result
	diff, err := git.GetDiff()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Check diff empty or not
	if diff == "" {
		fmt.Println("No staged changes found. Please stage your changes with 'git add'")
		os.Exit(1)
	}

	fmt.Println(diff)
}

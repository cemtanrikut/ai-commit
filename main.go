package main

import (
	"fmt"
	"os"

	"main.go/api"
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

	// Send diff to OpenAI API and generate commit message
	commitMessage, err := api.GenerateCommitMessage(diff)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Show the generated commit message
	fmt.Println("Generated Commit Message:")
	fmt.Println(commitMessage)

}

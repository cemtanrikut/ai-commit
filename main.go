package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	// Approvement check
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nDo you approve this commit message and commit changes? (y/n)")
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(strings.ToLower(answer))

	if answer == "y" || answer == "yes" {
		// Run commit process
		err := git.RunGitCommit(commitMessage)
		if err != nil {
			fmt.Println("Failed to commit: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Changes committed successfully!")
	} else {
		// Open an editor for edit commit message
		editedMessage, err := git.EditCommitMessage(commitMessage)
		if err != nil {
			fmt.Printf("Error editing message: %v\n", err)
			os.Exit(1)
		}

		// Commit when user complete edit
		err = git.RunGitCommit(editedMessage)
		if err != nil {
			fmt.Printf("Failed to commit: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Changes committed successfully!")
	}

}

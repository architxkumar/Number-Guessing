package helper

import "fmt"

// PrintGreetingMessage prints a welcome message amd basic rules.
func PrintGreetingMessage() {
	fmt.Println("Welcome to Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have limited number of chances to guess the correct number depending upon the difficulty level.")
	fmt.Println()
}

// PrintDifficultyLevel prints difficulty level with available chances.
func PrintDifficultyLevel() {
	fmt.Println("Please select the difficulty level.")
	fmt.Printf("1. Easy (%d)\n", easyLevelChances)
	fmt.Printf("2. Medium (%d)\n", mediumLevelChances)
	fmt.Printf("3. Hard (%d)\n", hardLevelChances)
	fmt.Println()
}

// PrintGameStartMessage prints user selected level and game start greeting.
func PrintGameStartMessage(selectedLevel int) {
	fmt.Println()
	switch selectedLevel {
	case easyLevel:
		fmt.Println("Great! You have selected the Easy difficulty level.")
	case mediumLevel:
		fmt.Println("Great! You have selected the Medium level.")
	case hardLevel:
		fmt.Println("Great! You have selected the Hard level.")
	}
	fmt.Println("Let's start the game!")
	fmt.Println()
}

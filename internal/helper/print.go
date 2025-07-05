package helper

import "fmt"

func PrintGreetingMessage() {
	fmt.Println("Welcome to Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")
	fmt.Println()
}

func PrintDifficultyLevel() {
	fmt.Println("Please select the difficulty level.")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Println()
}

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

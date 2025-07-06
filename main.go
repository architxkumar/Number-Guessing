package main

import (
	"Number-Guessing/internal/helper"
	"fmt"
	"strconv"
)

func main() {
	helper.PrintGreetingMessage()
	helper.PrintDifficultyLevel()
	var userInputString string
	var selectedLevel int

	for {
		fmt.Printf("Enter your choice: ")
		_, err := fmt.Scanln(&userInputString)
		if err != nil {
			fmt.Println("Error reading input, Try again")
			continue
		}
		selectedLevel, err = helper.DifficultyLevelValidation(userInputString)
		if err != nil {
			fmt.Println("Invalid choice, Try again")
			continue
		}
		break
	}
	helper.PrintGameStartMessage(selectedLevel)
	randomNumber := helper.RandomNumberGenerator()
	attempts := 0
	chances := helper.GetChancesForLevel(selectedLevel)
	for {
		fmt.Println()
		fmt.Printf("Enter your guess: ")
		_, err := fmt.Scanln(&userInputString)
		if err != nil {
			fmt.Println("Error reading input, Try again")
			continue
		}
		userInputNumber, err := strconv.Atoi(userInputString)
		if err != nil {
			fmt.Println("Error reading input, Try again")
			continue
		}
		attempts++
		chances--
		if userInputNumber < 1 || userInputNumber > 100 {
			fmt.Println("Invalid choice, Enter a number between 1 and 100")
			continue
		}
		if userInputNumber == randomNumber {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts\n", attempts)
			break
		}
		if chances > 0 {

			if randomNumber < userInputNumber {
				fmt.Printf("Incorrect! The number is less than %d\n", userInputNumber)
			} else {
				fmt.Printf("Incorrect! The number is greter than %d\n", userInputNumber)
			}
			continue
		} else {
			fmt.Println("Game Over! You have used all the attempts")
		}
		break
	}
}

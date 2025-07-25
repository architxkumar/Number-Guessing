package main

import (
	"Number-Guessing/internal/helper"
	"fmt"
	"strconv"
	"time"
)

func main() {
	helper.PrintGreetingMessage()
	for {
		helper.PrintDifficultyLevel()
		var userInput string
		var selectedLevel int

		for {
			fmt.Printf("Enter your choice: ")
			_, err := fmt.Scanln(&userInput)
			if err != nil {
				fmt.Println("Error reading input, Try again")
				continue
			}
			selectedLevel, err = helper.DifficultyLevelValidation(userInput)
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
		startTime := time.Now()
		for {
			fmt.Println()
			fmt.Printf("Enter your guess: ")
			_, err := fmt.Scanln(&userInput)
			if err != nil {
				fmt.Println("Error reading input, Try again")
				continue
			}
			guessedNumber, err := strconv.Atoi(userInput)
			if err != nil {
				fmt.Println("Error reading input, Try again")
				continue
			}
			// User is not penalize for system errors
			attempts++
			chances--
			// User is penalized for entering out-of-bounds input
			if guessedNumber < 1 || guessedNumber > 100 {
				fmt.Println("Invalid choice, Enter a number between 1 and 100")
				continue
			}
			if guessedNumber == randomNumber {
				fmt.Printf("Congratulations! You guessed the correct number in %d attempts\n", attempts)
				fmt.Printf("You took %v time\n", time.Since(startTime).Truncate(time.Second))
				break
			}
			if chances > 0 {
				if randomNumber < guessedNumber {
					fmt.Printf("Incorrect! The number is less than %d\n", guessedNumber)
				} else {
					fmt.Printf("Incorrect! The number is greater than %d\n", guessedNumber)
				}
				continue
			} else {
				fmt.Println("Game Over! You have used all the attempts")
				fmt.Printf("You took %v time\n", time.Since(startTime).Truncate(time.Second))
			}
			break
		}
		fmt.Println("")
		fmt.Printf("Would you like to play again? (Y/n): ")
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			fmt.Println("Error reading input, Try again")
			continue
		}
		if userInput == "y" || userInput == "Y" {
			continue
		}
		break
	}
}

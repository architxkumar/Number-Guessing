package helper

import (
	"errors"
	"strconv"
)

// DifficultyLevelValidation validates the user input and returns
// corresponding integer after conversion. It returns error in case
// of missing argument or invalid argument.
func DifficultyLevelValidation(userInput string) (int, error) {
	if userInput == "" {
		return 0, errors.New("arguments not supplied")
	}
	input, err := strconv.Atoi(userInput)
	if err != nil {
		return 0, err
	}
	if input < easyLevel || input > hardLevel {
		return 0, errors.New("invalid input")
	}
	return input, nil

}

// GetChancesForLevel returns the available chances for supplied level.
func GetChancesForLevel(level int) int {
	switch level {
	case easyLevel:
		return easyLevelChances
	case mediumLevel:
		return mediumLevelChances
	case hardLevel:
		return hardLevelChances
	}
	return 10
}

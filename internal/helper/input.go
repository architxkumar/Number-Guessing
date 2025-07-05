package helper

import (
	"errors"
	"strconv"
)

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

func GetInputsForLevel(level int) int {
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

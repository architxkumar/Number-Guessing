# Number Guessing Game

Number Guessing is a CLI game where the computer randomly selects a number 
and the user has to guess it. The user is given a limited number of chances to guess the number. 
If the user guesses the number correctly, the game ends, and the user wins. Otherwise,
the game continues until the user runs out of chances.

This is a learning project built in Go as part of [roadmap.sh](https://roadmap.sh/projects/number-guessing-game) challenge.

## Features

- Multiple difficulty levels (easy, medium, hard)
- Input validation
- Multiple rounds support

## Requirements

- Go 1.19 or later

## Installation

Run the following commands

```bash
git clone https://github.com/architxkumar/Number-Guessing.git
cd Number-Guessing
go build
```

## Usage
```bash
./Number-Guessing
```
### Sample Output 1 - Single Round
```text
Welcome to Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have limited number of chances to guess the correct number depending upon the difficulty level.

Please select the difficulty level.
1. Easy (10)
2. Medium (5)
3. Hard (3)

Enter your choice: 2

Great! You have selected the Medium level.
Let's start the game!


Enter your guess: 27
Incorrect! The number is greater than 27

Enter your guess: 60
Incorrect! The number is greater than 60

Enter your guess: 90
Incorrect! The number is less than 90

Enter your guess: 85
Incorrect! The number is less than 85

Enter your guess: 87
Game Over! You have used all the attempts
```
### Sample Output 2 - Multiple Rounds
```text
Welcome to Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have limited number of chances to guess the correct number depending upon the difficulty level.

Please select the difficulty level.
1. Easy (10)
2. Medium (5)
3. Hard (3)

Enter your choice: 2

Great! You have selected the Medium level.
Let's start the game!


Enter your guess: 60
Incorrect! The number is greater than 60

Enter your guess: 71
Incorrect! The number is less than 71

Enter your guess: 69
Incorrect! The number is less than 69

Enter your guess: 67
Incorrect! The number is less than 67

Enter your guess: 64
Congratulations! You guessed the correct number in 5 attempts

Would you like to play again? (Y/n): Y
Please select the difficulty level.
1. Easy (10)
2. Medium (5)
3. Hard (3)

Enter your choice: 3

Great! You have selected the Hard level.
Let's start the game!


Enter your guess: 75
Incorrect! The number is less than 75

Enter your guess: 38
Incorrect! The number is greater than 38

Enter your guess: 49
Game Over! You have used all the attempts

Would you like to play again? (Y/n): n
```

## License

This project is licensed under the [MIT License](./LICENSE)
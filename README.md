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
- Gameplay time measurement

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
## Example
```text
Welcome to Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have limited number of chances to guess the correct number depending upon the difficulty level.

Please select the difficulty level.
1. Easy (10)
2. Medium (5)
3. Hard (3)

Enter your choice: 1

Great! You have selected the Easy difficulty level.
Let's start the game!


Enter your guess: 29
Incorrect! The number is greater than 29

Enter your guess: 40
Incorrect! The number is greater than 40

Enter your guess: 60
Incorrect! The number is greater than 60

Enter your guess: 80
Incorrect! The number is greater than 80

Enter your guess: 100
Incorrect! The number is less than 100

Enter your guess: 99
Incorrect! The number is less than 99

Enter your guess: 95
Incorrect! The number is less than 95

Enter your guess: 89
Incorrect! The number is less than 89

Enter your guess: 84
Incorrect! The number is greater than 84

Enter your guess: 85
Game Over! You have used all the attempts
You took 27s time

Would you like to play again? (Y/n): n
```

## License

This project is licensed under the [MIT License](./LICENSE)
# Number Guessing Game

Number Guessing is a CLI game where the computer randomly selects a number 
and the user has to guess it. The user is given a limited number of chances to guess the number. 
If the user guesses the number correctly, the game ens, and the user wins. Otherwise,
the game continues until the user runs out of chances.
</br>
</br>
This is a learning project built in Go as part of [roadmap.sh](https://roadmap.sh/projects/number-guessing-game) challenge.

## Features

- Multiple Difficulty levels (easy, medium, hard)
- Input Validation

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
### Sample Output
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

## License

This project is licensed under the [MIT License](./LICENSE)
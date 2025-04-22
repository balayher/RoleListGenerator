package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Gets user input for a number, defaulting to 0 if an integer is not entered.
func getInput() (int, error) {
	userInput := bufio.NewReader(os.Stdin)
	userVal, err := userInput.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input := strings.TrimSpace(userVal)
	intVal, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return intVal, nil
}

// Gets user input for a yes/no response, defaulting to yes if response is invalid or empty.
func getYesNo() bool {
	userInput := bufio.NewReader(os.Stdin)
	userVal, _ := userInput.ReadString('\n')

	input := strings.TrimSpace(userVal)
	if input == "" {
		return true
	}

	input = strings.ToLower(input)
	if input == "y" || input == "yes" {
		return true
	}
	if input == "n" || input == "no" {
		return false
	}

	fmt.Println("Invalid input, value set to 'Yes'")
	return true
}

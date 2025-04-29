package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"slices"
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

// Gets user input for a yes/no response, defaulting to its default value if response is invalid or empty.
func getYesNo(defVal bool) bool {
	userInput := bufio.NewReader(os.Stdin)
	userVal, _ := userInput.ReadString('\n')

	input := strings.TrimSpace(userVal)
	if input == "" {
		fmt.Printf("No input found, value set to '%v'\n", defVal)
		return defVal
	}

	input = strings.ToLower(input)
	if input == "y" || input == "yes" {
		return true
	}
	if input == "n" || input == "no" {
		return false
	}

	fmt.Printf("Invalid input, value set to '%v'\n", defVal)
	return defVal
}

// Gets the name of a json file for quick input
func getJsonName() (string, bool) {
	userInput := bufio.NewReader(os.Stdin)
	userVal, _ := userInput.ReadString('\n')

	input := strings.TrimSpace(userVal)
	if input == "" {
		fmt.Println("No filename found")
		return "", false
	}

	if strings.HasSuffix(input, ".json") {
		return input, true
	}

	return input + ".json", true
}

// Gets user input for a list of roles banned from generation, defaulting to none.
// Formats the submitted roles to remove white space and ignore capitalization.
func getBanInput() ([]string, error) {
	userInput := bufio.NewReader(os.Stdin)
	userVal, err := userInput.ReadString('\n')
	if err != nil {
		return []string{}, err
	}

	input := strings.Split(strings.ToLower(userVal), " ")
	for i := range input {
		input[i] = strings.TrimSpace(input[i])
	}

	return input, nil
}

// Formats the role list output to make it more readable
// If numbered output is toggled, adds a distinct number to each role when printing

func formatOutput(roleList []string, roleNumbers []int, numbered bool) []int {
	for i := range roleList {
		if i > 0 {
			if i%5 == 0 {
				fmt.Println()
			} else {
				fmt.Print(", ")
			}
		}
		if numbered {
			randomIdx := rand.Intn(len(roleNumbers))
			fmt.Printf("%v.) ", roleNumbers[randomIdx])
			roleNumbers = slices.Delete(roleNumbers, randomIdx, randomIdx+1)
		}
		fmt.Printf("%v", roleList[i])
	}
	fmt.Println()
	if len(roleList) > 0 {
		fmt.Println()
	}
	return roleNumbers
}

func fileOutput(roleList []string, roleNumbers []int, numbered bool, f io.Writer) []int {
	for i := range roleList {
		if numbered {
			randomIdx := rand.Intn(len(roleNumbers))
			fmt.Fprintf(f, "%v.) ", roleNumbers[randomIdx])
			roleNumbers = slices.Delete(roleNumbers, randomIdx, randomIdx+1)
		}
		fmt.Fprintf(f, "%v\n", roleList[i])
	}
	fmt.Fprintln(f, "")
	return roleNumbers
}

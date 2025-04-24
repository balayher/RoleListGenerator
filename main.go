package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initializing input variables.
	var ti, tp, ts, tk, rt, mk, ms, md, rm, ce, nk, nc, ne, nb, rn, a, vamp int
	jailor, gf, cl, anyMaf, anyCov, anyVamp, custom := false, false, false, true, true, true, true
	var ban []string

	// Seeding randomization
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Get user input for the amount of each role category is requested.
	// If the input is not an integer, the value is defaulted to 0.
	// Prompts if user wants a guaranteed Jailor, Godfather, or Coven Leader if role could be generated.
	// If a role category has a maximum number, the value will be set to that maximum if exceeded.
	fmt.Print("Enter the number of Town Investigative: ")
	ti, err := getInput()
	if err != nil {
		fmt.Println("Invalid input, TI set to 0")
	}

	fmt.Print("Enter the number of Town Protective: ")
	tp, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, TP set to 0")
	}

	fmt.Print("Enter the number of Town Support: ")
	ts, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, TS set to 0")
	}

	fmt.Print("Enter the number of Town Killing: ")
	tk, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, TK set to 0")
	}
	fmt.Print("Enter the number of Random Town: ")
	rt, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, RT set to 0")
	}
	if (tk > 0) || (rt > 0) {
		fmt.Print("Do you want a guaranteed Jailor? ")
		jailor = getYesNo(jailor)
	}

	fmt.Print("Enter the number of Mafia Killing: ")
	mk, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, MK set to 0")
	}
	fmt.Print("Enter the number of Mafia Support: ")
	ms, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, MS set to 0")
	}
	fmt.Print("Enter the number of Mafia Deception: ")
	md, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, MD set to 0")
	}
	fmt.Print("Enter the number of Random Mafia: ")
	rm, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, RM set to 0")
	}
	if (mk > 0) || (rm > 0) {
		fmt.Print("Do you want a guaranteed Godfather? ")
		gf = getYesNo(gf)
	}

	if md+ms > 0 && mk+rm == 0 {
		mk++
		fmt.Println("MS or MD detected without MK or RM, one MK added.")
	}

	fmt.Print("Enter the number of Coven Evil: ")
	ce, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, CE set to 0")
	}
	if ce > 0 {
		fmt.Print("Do you want a guaranteed Coven Leader? ")
		cl = getYesNo(cl)
	}

	fmt.Print("Enter the number of Vampires: ")
	vamp, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, Vampires set to 0")
	}

	fmt.Print("Enter the number of Neutral Killing: ")
	nk, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, NK set to 0")
	}
	fmt.Print("Enter the number of Neutral Chaos: ")
	nc, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, NC set to 0")
	}
	fmt.Print("Enter the number of Neutral Evil: ")
	ne, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, NE set to 0")
	}
	fmt.Print("Enter the number of Neutral Benign: ")
	nb, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, NB set to 0")
	}
	fmt.Print("Enter the number of Random Netural: ")
	rn, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, RN set to 0")
	}

	fmt.Print("Enter the number of Any: ")
	a, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, Any set to 0")
	}

	if nc > 0 || rn > 0 || a > 0 {
		fmt.Print("Do you want Vampires as a random option? ")
		anyVamp = getYesNo(anyVamp)
	}

	// Prompts if Mafia or Coven should be available in Any if they don't already exist
	if a > 0 && (mk+ms+md+rm == 0) {
		fmt.Print("Do you want Mafia possible in your Any? ")
		anyMaf = getYesNo(anyMaf)
	}
	if a > 0 && ce == 0 {
		fmt.Print("Do you want Coven possible in your Any? This removes Witch from the pool. ")
		anyCov = getYesNo(anyCov)
	}

	fmt.Print("Do you want to use custom roles? ")
	custom = getYesNo(custom)

	fmt.Print("Do you want to ban any roles? Separate roles with a space, and use a _ for any multiple word role (such as Coven_Leader):\n")
	ban, err = getBanInput()
	if err != nil {
		fmt.Println("Invalid input, no ban list set")
	}
	fmt.Println()

	// Calls createRoles to generate each set of roles, then prints them to terminal.
	town, mafia, coven, neutral, anyRole := createRoles(ti, tp, ts, tk, rt, mk, ms, md, rm, ce, nk, nc, ne, nb, rn, a, vamp, jailor, gf, cl, anyMaf, anyCov, anyVamp, custom, ban)
	fmt.Println()
	fmt.Println("Town:")
	formatOutput(town)
	fmt.Println("Mafia:")
	formatOutput(mafia)
	fmt.Println("Coven:")
	formatOutput(coven)
	fmt.Println("Neutral:")
	formatOutput(neutral)
	fmt.Println("Any:")
	formatOutput(anyRole)
}

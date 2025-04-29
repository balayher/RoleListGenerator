package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Initializing input variables.
	var ban []string
	var roleNumbers []int

	// Seeding randomization
	rand.New(rand.NewSource(time.Now().UnixNano()))

	c := getCounts(false)
	t := getToggles(c, false)

	// Allows for preventing specific roles from appearing in the role list.
	// This option may reduce the size of the final list if there are no roles left to generate in a requested category
	fmt.Print("Do you want to ban any roles? Separate roles with a space, and use a _ for any multiple word role (such as 'Coven_Leader'). Turncoat will need a specified faction (i.e. 'Turncoat(Mafia)'):\n")
	ban, err := getBanInput()
	if err != nil {
		fmt.Println("Invalid input, no ban list set")
	}

	fmt.Println()

	// Calls createRoles to generate all of the roles.
	rl := createRoles(c, t, ban)
	totalRoles := len(rl.town) + len(rl.mafia) + len(rl.coven) + len(rl.neutral) + len(rl.allAny)

	// Setup for numbering roles if option is on
	if t.numbered {
		for i := 1; i < totalRoles+1; i++ {
			roleNumbers = append(roleNumbers, i)
		}
	}

	fmt.Println()
	fmt.Printf("%v roles generated.\n\n", totalRoles)

	// Writes the roles to roles.txt if option is selected.
	// Otherwise prints the roles to the terminal.
	if t.fileWrite {
		f, err := os.Create("roles.txt")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		if len(rl.town) > 0 {
			fmt.Fprintln(f, "Town:")
			roleNumbers = fileOutput(rl.town, roleNumbers, t.numbered, f)
		}
		if len(rl.mafia) > 0 {
			fmt.Fprintln(f, "Mafia:")
			roleNumbers = fileOutput(rl.mafia, roleNumbers, t.numbered, f)
		}
		if len(rl.coven) > 0 {
			fmt.Fprintln(f, "Coven:")
			roleNumbers = fileOutput(rl.coven, roleNumbers, t.numbered, f)
		}
		if len(rl.neutral) > 0 {
			fmt.Fprintln(f, "Neutral:")
			roleNumbers = fileOutput(rl.neutral, roleNumbers, t.numbered, f)
		}
		if len(rl.allAny) > 0 {
			fmt.Fprintln(f, "Any:")
			_ = fileOutput(rl.allAny, roleNumbers, t.numbered, f)
		}
		if len(rl.exeTargets) > 0 {
			fmt.Fprintln(f, "Executioner Targets:")
			_ = fileOutput(rl.exeTargets, roleNumbers, false, f)
		}
		if len(rl.gaTargets) > 0 {
			fmt.Fprintln(f, "Guardian Angel Targets:")
			_ = fileOutput(rl.gaTargets, roleNumbers, false, f)
		}

		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("roles.txt written successfully")
	} else {
		if len(rl.town) > 0 {
			fmt.Println("Town:")
			roleNumbers = formatOutput(rl.town, roleNumbers, t.numbered)
		}
		if len(rl.mafia) > 0 {
			fmt.Println("Mafia:")
			roleNumbers = formatOutput(rl.mafia, roleNumbers, t.numbered)
		}
		if len(rl.coven) > 0 {
			fmt.Println("Coven:")
			roleNumbers = formatOutput(rl.coven, roleNumbers, t.numbered)
		}
		if len(rl.neutral) > 0 {
			fmt.Println("Neutral:")
			roleNumbers = formatOutput(rl.neutral, roleNumbers, t.numbered)
		}
		if len(rl.allAny) > 0 {
			fmt.Println("Any:")
			_ = formatOutput(rl.allAny, roleNumbers, t.numbered)
		}
		if len(rl.exeTargets) > 0 {
			fmt.Println("Executioner Targets:")
			_ = formatOutput(rl.exeTargets, roleNumbers, false)
		}
		if len(rl.gaTargets) > 0 {
			fmt.Println("Guardian Angel Targets:")
			_ = formatOutput(rl.gaTargets, roleNumbers, false)
		}
	}
}

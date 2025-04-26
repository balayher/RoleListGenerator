package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Initializing input variables.
	var ti, tp, ts, tk, rt, mk, ms, md, rm, ce, nk, nc, ne, nb, rn, a, vamp int
	jailor, gf, cl, anyMaf, anyCov, anyVamp, custom, numbered, fileWrite := false, false, false, true, true, true, true, true, false
	var ban []string
	var roleNumbers []int

	// Seeding randomization
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Get user input for the amount of each role category is requested.
	// If the input is not an integer, the value is set to 0.
	// Prompts if user wants a guaranteed Jailor, Godfather, or Coven Leader if the role could be generated.
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

	// Checks if Mafia Support or Deception is added without Mafia Killing or Random Mafia
	// If so, one is changed to Mafia Killing to generate a Godfather or Mafioso
	if ms > 0 && mk+rm == 0 {
		ms--
		mk++
		fmt.Println("MS detected without MK or RM, replacing one MS with MK")
	} else if md > 0 && mk+rm == 0 {
		md--
		mk++
		fmt.Println("MD detected without MK or RM, replacing one MD with MK.")
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

	// Prompts if Vampires can be randomly generated
	// This allows user to add guaranteed Vampires without having more appear randomly
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

	// Toggle whether user wants to add the ISFL server custom roles or just have vanilla Town of Salem roles only
	fmt.Print("Do you want to use custom roles? ")
	custom = getYesNo(custom)

	// Allows for preventing specific roles from appearing in the role list.
	// This option may reduce the size of the final list if there are no roles left to generate in a requested category
	fmt.Print("Do you want to ban any roles? Separate roles with a space, and use a _ for any multiple word role (such as 'Coven_Leader'). Turncoat will need a specified faction (i.e. 'Turncoat(Mafia)'):\n")
	ban, err = getBanInput()
	if err != nil {
		fmt.Println("Invalid input, no ban list set")
	}

	// Toggle whether the roles are numbered in the output
	fmt.Print("Would you like to randomly number the roles for easy assignment? ")
	numbered = getYesNo(numbered)

	// Toggle whether the output is printed to terminal or to the roles.txt file
	fmt.Print("Would you like your rolelist written to a roles.txt file? ")
	fileWrite = getYesNo(fileWrite)

	fmt.Println()

	// Calls createRoles to generate all of the roles.
	town, mafia, coven, neutral, anyRole, exeTargets, gaTargets := createRoles(ti, tp, ts, tk, rt, mk, ms, md, rm, ce, nk, nc, ne, nb, rn, a, vamp, jailor, gf, cl, anyMaf, anyCov, anyVamp, custom, ban)
	totalRoles := len(town) + len(mafia) + len(coven) + len(neutral) + len(anyRole)

	// Setup for numbering roles if option is on
	if numbered {
		for i := 1; i < totalRoles+1; i++ {
			roleNumbers = append(roleNumbers, i)
		}
	}

	fmt.Println()
	fmt.Printf("%v roles generated.\n\n", totalRoles)

	// Writes the roles to roles.txt if option is selected.
	// Otherwise prints the roles to the terminal.
	if fileWrite {
		f, err := os.Create("roles.txt")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		if len(town) > 0 {
			fmt.Fprintln(f, "Town:")
			roleNumbers = fileOutput(town, roleNumbers, numbered, f)
		}
		if len(mafia) > 0 {
			fmt.Fprintln(f, "Mafia:")
			roleNumbers = fileOutput(mafia, roleNumbers, numbered, f)
		}
		if len(coven) > 0 {
			fmt.Fprintln(f, "Coven:")
			roleNumbers = fileOutput(coven, roleNumbers, numbered, f)
		}
		if len(neutral) > 0 {
			fmt.Fprintln(f, "Neutral:")
			roleNumbers = fileOutput(neutral, roleNumbers, numbered, f)
		}
		if len(anyRole) > 0 {
			fmt.Fprintln(f, "Any:")
			_ = fileOutput(anyRole, roleNumbers, numbered, f)
		}
		if len(exeTargets) > 0 {
			fmt.Fprintln(f, "Executioner Targets:")
			_ = fileOutput(exeTargets, roleNumbers, false, f)
		}
		if len(gaTargets) > 0 {
			fmt.Fprintln(f, "Guardian Angel Targets:")
			_ = fileOutput(gaTargets, roleNumbers, false, f)
		}

		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("roles.txt written successfully")
	} else {
		if len(town) > 0 {
			fmt.Println("Town:")
			roleNumbers = formatOutput(town, roleNumbers, numbered)
		}
		if len(mafia) > 0 {
			fmt.Println("Mafia:")
			roleNumbers = formatOutput(mafia, roleNumbers, numbered)
		}
		if len(coven) > 0 {
			fmt.Println("Coven:")
			roleNumbers = formatOutput(coven, roleNumbers, numbered)
		}
		if len(neutral) > 0 {
			fmt.Println("Neutral:")
			roleNumbers = formatOutput(neutral, roleNumbers, numbered)
		}
		if len(anyRole) > 0 {
			fmt.Println("Any:")
			_ = formatOutput(anyRole, roleNumbers, numbered)
		}
		if len(exeTargets) > 0 {
			fmt.Println("Executioner Targets:")
			_ = formatOutput(exeTargets, roleNumbers, false)
		}
		if len(gaTargets) > 0 {
			fmt.Println("Guardian Angel Targets:")
			_ = formatOutput(gaTargets, roleNumbers, false)
		}
	}
}

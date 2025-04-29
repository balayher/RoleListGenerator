package main

import "fmt"

type Counts struct {
	ti   int `json:"ti"`
	tp   int `json:"tp"`
	ts   int `json:"ts"`
	tk   int `json:"tk"`
	rt   int `json:"rt"`
	mk   int `json:"mk"`
	ms   int `json:"ms"`
	md   int `json:"md"`
	rm   int `json:"rm"`
	ce   int `json:"ce"`
	nk   int `json:"nk"`
	nc   int `json:"nc"`
	ne   int `json:"ne"`
	nb   int `json:"nb"`
	rn   int `json:"rn"`
	a    int `json:"a"`
	vamp int `json:"vamp"`
}

type Toggles struct {
	jailor    bool `json:"jailor"`
	gf        bool `json:"gf"`
	cl        bool `json:"cl"`
	anyMaf    bool `json:"anyMaf"`
	anyCov    bool `json:"anyCov"`
	anyVamp   bool `json:"anyVamp"`
	custom    bool `json:"custom"`
	numbered  bool `json:"numbered"`
	fileWrite bool `json:"fileWrite"`
}

type RoleList struct {
	town       []string
	mafia      []string
	coven      []string
	neutral    []string
	allAny     []string
	exeList    []string
	gaList     []string
	exeTargets []string
	gaTargets  []string
}

func getCounts(counts bool) Counts {
	c := Counts{
		ti:   0,
		tp:   0,
		ts:   0,
		tk:   0,
		rt:   0,
		mk:   0,
		ms:   0,
		md:   0,
		rm:   0,
		ce:   0,
		nk:   0,
		nc:   0,
		ne:   0,
		nb:   0,
		rn:   0,
		a:    0,
		vamp: 0,
	}
	var err error

	if !counts {
		// Get user input for the amount of each role category is requested.
		// If the input is not an integer, the value is set to 0.
		// Prompts if user wants a guaranteed Jailor, Godfather, or Coven Leader if the role could be generated.
		fmt.Print("Enter the number of Town Investigative: ")
		c.ti, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, TI set to 0")
		}
		fmt.Print("Enter the number of Town Protective: ")
		c.tp, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, TP set to 0")
		}
		fmt.Print("Enter the number of Town Support: ")
		c.ts, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, TS set to 0")
		}
		fmt.Print("Enter the number of Town Killing: ")
		c.tk, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, TK set to 0")
		}
		fmt.Print("Enter the number of Random Town: ")
		c.rt, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, RT set to 0")
		}

		fmt.Print("Enter the number of Mafia Killing: ")
		c.mk, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, MK set to 0")
		}
		fmt.Print("Enter the number of Mafia Support: ")
		c.ms, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, MS set to 0")
		}
		fmt.Print("Enter the number of Mafia Deception: ")
		c.md, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, MD set to 0")
		}
		fmt.Print("Enter the number of Random Mafia: ")
		c.rm, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, RM set to 0")
		}

		// Checks if Mafia Support or Deception is added without Mafia Killing or Random Mafia
		// If so, one is changed to Mafia Killing to generate a Godfather or Mafioso
		if c.ms > 0 && c.mk+c.rm == 0 {
			c.ms--
			c.mk++
			fmt.Println("MS detected without MK or RM, replacing one MS with MK")
		} else if c.md > 0 && c.mk+c.rm == 0 {
			c.md--
			c.mk++
			fmt.Println("MD detected without MK or RM, replacing one MD with MK.")
		}

		fmt.Print("Enter the number of Coven Evil: ")
		c.ce, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, CE set to 0")
		}

		fmt.Print("Enter the number of Vampires: ")
		c.vamp, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, Vampires set to 0")
		}

		fmt.Print("Enter the number of Neutral Killing: ")
		c.nk, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, NK set to 0")
		}
		fmt.Print("Enter the number of Neutral Chaos: ")
		c.nc, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, NC set to 0")
		}
		fmt.Print("Enter the number of Neutral Evil: ")
		c.ne, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, NE set to 0")
		}
		fmt.Print("Enter the number of Neutral Benign: ")
		c.nb, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, NB set to 0")
		}
		fmt.Print("Enter the number of Random Netural: ")
		c.rn, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, RN set to 0")
		}

		fmt.Print("Enter the number of Any: ")
		c.a, err = getInput()
		if err != nil {
			fmt.Println("Invalid input, Any set to 0")
		}
		return c
	}
	return c
}

func getToggles(c Counts, toggles bool) Toggles {
	t := Toggles{
		jailor:    false,
		gf:        false,
		cl:        false,
		anyMaf:    true,
		anyCov:    true,
		anyVamp:   true,
		custom:    true,
		numbered:  true,
		fileWrite: false,
	}

	if !toggles {
		if (c.tk > 0) || (c.rt > 0) {
			fmt.Print("Do you want a guaranteed Jailor? ")
			t.jailor = getYesNo(t.jailor)
		}
		if (c.mk > 0) || (c.rm > 0) {
			fmt.Print("Do you want a guaranteed Godfather? ")
			t.gf = getYesNo(t.gf)
		}
		if c.ce > 0 {
			fmt.Print("Do you want a guaranteed Coven Leader? ")
			t.cl = getYesNo(t.cl)
		}

		// Prompts if Vampires can be randomly generated
		// This allows user to add guaranteed Vampires without having more appear randomly
		if c.nc > 0 || c.rn > 0 || c.a > 0 {
			fmt.Print("Do you want Vampires as a random option? ")
			t.anyVamp = getYesNo(t.anyVamp)
		}

		// Prompts if Mafia or Coven should be available in Any if they don't already exist
		if c.a > 0 && (c.mk+c.ms+c.md+c.rm == 0) {
			fmt.Print("Do you want Mafia possible in your Any? ")
			t.anyMaf = getYesNo(t.anyMaf)
		}
		if c.a > 0 && c.ce == 0 {
			fmt.Print("Do you want Coven possible in your Any? This removes Witch from the pool. ")
			t.anyCov = getYesNo(t.anyCov)
		}

		// Toggle whether user wants to add the ISFL server custom roles or just have vanilla Town of Salem roles only
		fmt.Print("Do you want to use custom roles? ")
		t.custom = getYesNo(t.custom)

		// Toggle whether the roles are numbered in the output
		fmt.Print("Would you like to randomly number the roles for easy assignment? ")
		t.numbered = getYesNo(t.numbered)

		// Toggle whether the output is printed to terminal or to the roles.txt file
		fmt.Print("Would you like your rolelist written to a roles.txt file? ")
		t.fileWrite = getYesNo(t.fileWrite)

	}

	return t
}

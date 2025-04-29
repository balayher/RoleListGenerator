package main

import (
	"fmt"
)

type Counts struct {
	TI   int `json:"ti"`
	TP   int `json:"tp"`
	TS   int `json:"ts"`
	TK   int `json:"tk"`
	RT   int `json:"rt"`
	MK   int `json:"mk"`
	MS   int `json:"ms"`
	MD   int `json:"md"`
	RM   int `json:"rm"`
	CE   int `json:"ce"`
	NK   int `json:"nk"`
	NC   int `json:"nc"`
	NE   int `json:"ne"`
	NB   int `json:"nb"`
	RN   int `json:"rn"`
	A    int `json:"a"`
	Vamp int `json:"vamp"`
}

type Options struct {
	Jailor    bool `json:"jailor"`
	GF        bool `json:"gf"`
	CL        bool `json:"cl"`
	AnyMaf    bool `json:"anyMaf"`
	AnyCov    bool `json:"anyCov"`
	AnyVamp   bool `json:"anyVamp"`
	Custom    bool `json:"custom"`
	Numbered  bool `json:"numbered"`
	FileWrite bool `json:"fileWrite"`
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

func getCounts() Counts {
	c := Counts{
		TI:   0,
		TP:   0,
		TS:   0,
		TK:   0,
		RT:   0,
		MK:   0,
		MS:   0,
		MD:   0,
		RM:   0,
		CE:   0,
		NK:   0,
		NC:   0,
		NE:   0,
		NB:   0,
		RN:   0,
		A:    0,
		Vamp: 0,
	}
	var err error

	fmt.Println("Do you want to upload role counts from a .json file?")
	counts := getYesNo(false)

	if counts {
		fmt.Println("Enter the name of your .json file")
		jsonCountsFile, validName := getJsonName()
		if validName {
			jsonCounts, validJson := getJsonCounts(c, jsonCountsFile)
			if validJson {
				return jsonCounts
			}
		}
	}

	// Get user input for the amount of each role category is requested.
	// If the input is not an integer, the value is set to 0.
	// Prompts if user wants a guaranteed Jailor, Godfather, or Coven Leader if the role could be generated.
	fmt.Print("Enter the number of Town Investigative: ")
	c.TI, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, TI set to 0")
	}
	fmt.Print("Enter the number of Town Protective: ")
	c.TP, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, TP set to 0")
	}
	fmt.Print("Enter the number of Town Support: ")
	c.TS, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, TS set to 0")
	}
	fmt.Print("Enter the number of Town Killing: ")
	c.TK, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, TK set to 0")
	}
	fmt.Print("Enter the number of Random Town: ")
	c.RT, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, RT set to 0")
	}

	fmt.Print("Enter the number of Mafia Killing: ")
	c.MK, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, MK set to 0")
	}
	fmt.Print("Enter the number of Mafia Support: ")
	c.MS, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, MS set to 0")
	}
	fmt.Print("Enter the number of Mafia Deception: ")
	c.MD, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, MD set to 0")
	}
	fmt.Print("Enter the number of Random Mafia: ")
	c.RM, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, RM set to 0")
	}

	// Checks if Mafia Support or Deception is added without Mafia Killing or Random Mafia
	// If so, one is changed to Mafia Killing to generate a Godfather or Mafioso
	if c.MS > 0 && c.MK+c.RM == 0 {
		c.MS--
		c.MK++
		fmt.Println("MS detected without MK or RM, replacing one MS with MK")
	} else if c.MD > 0 && c.MK+c.RM == 0 {
		c.MD--
		c.MK++
		fmt.Println("MD detected without MK or RM, replacing one MD with MK.")
	}

	fmt.Print("Enter the number of Coven Evil: ")
	c.CE, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, CE set to 0")
	}

	fmt.Print("Enter the number of Vampires: ")
	c.Vamp, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, Vampires set to 0")
	}

	fmt.Print("Enter the number of Neutral Killing: ")
	c.NK, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, NK set to 0")
	}
	fmt.Print("Enter the number of Neutral Chaos: ")
	c.NC, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, NC set to 0")
	}
	fmt.Print("Enter the number of Neutral Evil: ")
	c.NE, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, NE set to 0")
	}
	fmt.Print("Enter the number of Neutral Benign: ")
	c.NB, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, NB set to 0")
	}
	fmt.Print("Enter the number of Random Netural: ")
	c.RN, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, RN set to 0")
	}

	fmt.Print("Enter the number of Any: ")
	c.A, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, Any set to 0")
	}

	fmt.Println("Do you want to save your role counts to counts.json?")
	saveCounts := getYesNo(false)

	if saveCounts {
		saveJson(c, "counts.json")
	}

	return c
}

func getOptions(c Counts) Options {
	t := Options{
		Jailor:    false,
		GF:        false,
		CL:        false,
		AnyMaf:    true,
		AnyCov:    true,
		AnyVamp:   true,
		Custom:    true,
		Numbered:  true,
		FileWrite: false,
	}

	fmt.Println("Do you want to upload role options from a .json file?")
	toggles := getYesNo(false)

	if toggles {
		fmt.Println("Enter the name of your .json file")
		jsonOptionsFile, validName := getJsonName()
		if validName {
			jsonOptions, validJson := getJsonOptions(t, jsonOptionsFile)
			if validJson {
				return jsonOptions
			}
		}
	}

	if (c.TK > 0) || (c.RT > 0) {
		fmt.Print("Do you want a guaranteed Jailor? ")
		t.Jailor = getYesNo(t.Jailor)
	}
	if (c.MK > 0) || (c.RM > 0) {
		fmt.Print("Do you want a guaranteed Godfather? ")
		t.GF = getYesNo(t.GF)
	}
	if c.CE > 0 {
		fmt.Print("Do you want a guaranteed Coven Leader? ")
		t.CL = getYesNo(t.CL)
	}

	// Prompts if Vampires can be randomly generated
	// This allows user to add guaranteed Vampires without having more appear randomly
	if c.NC > 0 || c.RN > 0 || c.A > 0 {
		fmt.Print("Do you want Vampires as a random option? ")
		t.AnyVamp = getYesNo(t.AnyVamp)
	}

	// Prompts if Mafia or Coven should be available in Any slots

	fmt.Print("Do you want Mafia possible in your Any? ")
	t.AnyMaf = getYesNo(t.AnyMaf)

	fmt.Print("Do you want Coven possible in your Any? This removes Witch from the pool. ")
	t.AnyCov = getYesNo(t.AnyCov)

	// Toggle whether user wants to add the ISFL server custom roles or just have vanilla Town of Salem roles only
	fmt.Print("Do you want to use custom roles? ")
	t.Custom = getYesNo(t.Custom)

	// Toggle whether the roles are numbered in the output
	fmt.Print("Would you like to randomly number the roles for easy assignment? ")
	t.Numbered = getYesNo(t.Numbered)

	// Toggle whether the output is printed to terminal or to the roles.txt file
	fmt.Print("Would you like your rolelist written to a roles.txt file? ")
	t.FileWrite = getYesNo(t.FileWrite)

	fmt.Println("Do you want to save your role counts to options.json?")
	saveOptions := getYesNo(false)

	if saveOptions {
		saveJson(t, "options.json")
	}

	return t
}

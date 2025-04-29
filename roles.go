package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

func createRoles(c Counts, t Options, ban []string) RoleList {
	// Initializes slices for each faction for the final role list.
	rl := RoleList{
		town:       []string{},
		mafia:      []string{},
		coven:      []string{},
		neutral:    []string{},
		allAny:     []string{},
		exeList:    []string{},
		gaList:     []string{},
		exeTargets: []string{},
		gaTargets:  []string{},
	}

	extra := 0

	// Defines each Town role category.
	townInvestigative := []string{
		"Investigator",
		"Sheriff",
		"Lookout",
		"Tracker",
		"Psychic",
		"Spy",
	}
	townProtective := []string{
		"Bodyguard",
		"Doctor",
		"Crusader",
		"Trapper",
	}
	townSupport := []string{
		"Mayor",
		"Escort",
		"Retributionist",
		"Medium",
		"Transporter",
	}
	townKilling := []string{
		"Jailor",
		"Veteran",
		"Vigilante",
	}

	// Defines each Mafia role category, plus a Godfather/Mafioso slice to guarantee one exists when necessary.
	gfMafioso := []string{
		"Godfather",
		"Mafioso",
	}
	mafiaKilling := []string{
		"Godfather",
		"Mafioso",
		"Ambusher",
	}
	mafiaSupport := []string{
		"Consort",
		"Blackmailer",
		"Consigliere",
	}
	mafiaDeception := []string{
		"Disguiser",
		"Forger",
		"Framer",
		"Hypnotist",
		"Janitor",
	}

	// Defines the Coven role category.
	covenEvil := []string{
		"Coven_Leader",
		"Hex_Master",
		"Medusa",
		"Potion_Master",
		"Necromancer",
		"Poisoner",
	}

	// Defines each Netural role category.
	neutralKilling := []string{
		"Arsonist",
		"Juggernaut",
		"Serial_Killer",
		"Werewolf",
	}
	neutralEvil := []string{
		"Executioner",
		"Jester",
		"Witch",
	}
	neutralChaos := []string{
		"Pirate",
		"Plaguebearer",
	}
	neutralBenign := []string{
		"Amnesiac",
		"Guardian_Angel",
		"Survivor",
	}

	// If custom roles are toggled on, adds the custom roles to the role category they belong to.
	if t.Custom {
		townInvestigative = append(townInvestigative, "Seer", "Detective")
		townProtective = append(townProtective, "Cleric", "Oracle")
		townSupport = append(townSupport, "Monarch", "Governor", "Prosecutor", "Jack_of_All_Trades", "Timeshifter")
		townKilling = append(townKilling, "Gambler")
		mafiaKilling = append(mafiaKilling, "Poppet")
		mafiaSupport = append(mafiaSupport, "Watcher", "Angler", "Underboss", "Bouncer")
		mafiaDeception = append(mafiaDeception, "Stager")
		covenEvil = append(covenEvil, "Soultaker", "Siren", "Voodoo_Queen", "Frostbringer")
		neutralKilling = append(neutralKilling, "Mutator", "Horticulturist", "Shapeshifter", "Shroud", "Bombardier", "Gargoyle")
		neutralChaos = append(neutralChaos, "Inquisitor", "Anarchist", "Quack", "Stalker")
		neutralEvil = append(neutralEvil, "Turncoat(Mafia)", "Turncoat(Coven)")
	}

	// Remove user requested banned roles from the list
	townInvestigative = checkBans(townInvestigative, ban)
	townProtective = checkBans(townProtective, ban)
	townSupport = checkBans(townSupport, ban)
	townKilling = checkBans(townKilling, ban)
	gfMafioso = checkBans(gfMafioso, ban)
	mafiaKilling = checkBans(mafiaKilling, ban)
	mafiaSupport = checkBans(mafiaSupport, ban)
	mafiaDeception = checkBans(mafiaDeception, ban)
	covenEvil = checkBans(covenEvil, ban)
	neutralKilling = checkBans(neutralKilling, ban)
	neutralEvil = checkBans(neutralEvil, ban)
	neutralChaos = checkBans(neutralChaos, ban)
	neutralBenign = checkBans(neutralBenign, ban)

	// Defines which roles are unique.
	unique := []string{
		"Cleric",
		"Oracle",
		"Mayor",
		"Retributionist",
		"Governor",
		"Prosecutor",
		"Monarch",
		"Jailor",
		"Veteran",
		"Godfather",
		"Mafioso",
		"Ambusher",
		"Angler",
		"Poppet",
		"Underboss",
		"Bouncer",
		"Coven_Leader",
		"Hex_Master",
		"Medusa",
		"Potion_Master",
		"Necromancer",
		"Poisoner",
		"Soultaker",
		"Siren",
		"Voodoo_Queen",
		"Frostbringer",
		"Juggernaut",
		"Werewolf",
		"Pirate",
		"Plaguebearer",
		"Mutator",
		"Inquisitor",
		"Anarchist",
	}

	// Defines which Town roles cannot be Executioner targets.
	nonExe := []string{
		"Jailor",
		"Mayor",
		"Governor",
		"Monarch",
		"Prosecutor",
	}

	// Defines which Unique Town roles can be Executioner targets (for use when assigning Any roles).
	uniqueExe := []string{
		"Cleric",
		"Oracle",
		"Retributionist",
		"Veteran",
	}

	// Defines which roles cannot be Guardian Angel targets.
	nonGA := []string{
		"Executioner",
		"Jester",
		"Guardian_Angel",
	}

	// Adds Vampires to random pool if allowed
	if t.AnyVamp && !slices.Contains(ban, "vampire") {
		neutralChaos = append(neutralChaos, "Vampire")
	}

	// Cancels guaranteed roles if they are on the ban list
	if slices.Contains(ban, "jailor") {
		t.Jailor = false
	}
	if slices.Contains(ban, "godfather") {
		t.GF = false
	}
	if slices.Contains(ban, "coven_leader") {
		t.CL = false
	}

	// Checks if both Godfather and Mafioso are on the ban list and converts all Mafia slots to Any slots
	if len(gfMafioso) == 0 {
		maf := c.MK + c.MS + c.MD + c.RM
		fmt.Printf("Both Godfather and Mafioso banned, converting %v mafia slots to Any.\n", maf)
		c.A += maf
		c.MK = 0
		c.MS = 0
		c.MD = 0
		c.RM = 0
		t.AnyMaf = false
	}

	// Adds Godfather if guaranteed, else adds either Godfather or Mafioso if Mafia exists.
	if t.GF && c.MK > 0 {
		fmt.Println("Adding Godfather.")
		c.MK, mafiaKilling, rl.mafia = insertGuaranteedRole(c.MK, mafiaKilling, rl.mafia, "Godfather")
	} else if t.GF && c.RM > 0 {
		fmt.Println("Adding Godfather.")
		c.RM, mafiaKilling, rl.mafia = insertGuaranteedRole(c.RM, mafiaKilling, rl.mafia, "Godfather")
	} else if c.MK > 0 {
		fmt.Println("Adding Godfather or Mafioso.")
		_, rl.mafia, _ = randomRoleSelection(1, gfMafioso, unique, rl.mafia)
		c.MK--
		mafiaKilling = removeUnique(rl.mafia[0], mafiaKilling)
	} else if c.RM > 0 {
		fmt.Println("Adding Godfather or Mafioso.")
		_, rl.mafia, _ = randomRoleSelection(1, gfMafioso, unique, rl.mafia)
		c.RM--
		mafiaKilling = removeUnique(rl.mafia[0], mafiaKilling)
	}

	// Converts Mafia subcategories to Random Mafia if no roles are available in the subcategory and skips role selection.
	if len(mafiaKilling) == 0 && c.MK > 0 {
		fmt.Printf("No valid Mafia Killing roles, %v slots converted to Random Mafia.\n", c.MK)
		c.RM += c.MK
		c.MK = 0
	}
	if len(mafiaDeception) == 0 && c.MD > 0 {
		fmt.Printf("No valid Mafia Deception roles, %v slots converted to Random Mafia.\n", c.MD)
		c.RM += c.MD
		c.MD = 0
	}
	if len(mafiaSupport) == 0 && c.MS > 0 {
		fmt.Printf("No valid Mafia Support roles, %v slots converted to Random Mafia.\n", c.MS)
		c.RM += c.MS
		c.MS = 0
	}

	// Adds all Mafia roles requested.
	if c.MK > 0 {
		fmt.Printf("Adding %v Mafia Killing.\n", c.MK)
		mafiaKilling, rl.mafia, extra = randomRoleSelection(c.MK, mafiaKilling, unique, rl.mafia)
		if extra > 0 {
			fmt.Printf("Converting %v slots to Random Mafia.\n", extra)
			c.RM += extra
		}
	}
	if c.MD > 0 {
		fmt.Printf("Adding %v Mafia Deception.\n", c.MD)
		mafiaDeception, rl.mafia, extra = randomRoleSelection(c.MD, mafiaDeception, unique, rl.mafia)
		if extra > 0 {
			fmt.Printf("Converting %v slots to Random Mafia.\n", extra)
			c.RM += extra
		}
	}
	if c.MS > 0 {
		fmt.Printf("Adding %v Mafia Support.\n", c.MS)
		mafiaSupport, rl.mafia, extra = randomRoleSelection(c.MS, mafiaSupport, unique, rl.mafia)
		if extra > 0 {
			fmt.Printf("Converting %v slots to Random Mafia.\n", extra)
			c.RM += extra
		}
	}

	randomMafia := slices.Concat(mafiaKilling, mafiaDeception, mafiaSupport)
	if len(randomMafia) == 0 && c.RM > 0 {
		fmt.Printf("No valid Mafia roles, %v slots converted to Any.\n", c.RM)
		c.A += c.RM
		c.RM = 0
	}
	if c.RM > 0 {
		fmt.Printf("Adding %v Random Mafia.\n", c.RM)
		randomMafia, rl.mafia, extra = randomRoleSelection(c.RM, randomMafia, unique, rl.mafia)
		if extra > 0 {
			fmt.Printf("Converting %v slots to Any.\n", extra)
			c.A += extra
		}
	}

	// Adds rolled Mafia roles to Guardian Angel target list.
	if len(rl.mafia) > 0 {
		rl.gaList = append(rl.gaList, rl.mafia...)
	}

	// Converts Coven slots to Any if all Coven roles are banned.
	if len(covenEvil) == 0 && c.CE > 0 {
		fmt.Printf("No valid Coven roles, %v slots converted to Any.\n", c.CE)
		c.A += c.CE
		c.CE = 0
	}

	// Adds Coven Leader if guaranteed, then adds all Coven roles requested.
	if t.CL && c.CE > 0 {
		fmt.Println("Adding Coven Leader.")
		c.CE, covenEvil, rl.coven = insertGuaranteedRole(c.CE, covenEvil, rl.coven, "Coven_Leader")
	}
	if c.CE > 0 {
		fmt.Printf("Adding %v Coven Evil.\n", c.CE)
		covenEvil, rl.coven, extra = randomRoleSelection(c.CE, covenEvil, unique, rl.coven)
		if extra > 0 {
			fmt.Printf("Converting %v slots to Any.\n", extra)
			c.A += extra
		}
	}

	// Adds rolled Coven roles to Guardian Angel target list.
	if len(rl.coven) > 0 {
		rl.gaList = append(rl.gaList, rl.coven...)
	}

	// Removes Turncoats from the NE list if either Mafia or Coven doesn't exist.
	// Removes Witch from the NE list if Coven exists or can be rolled in an Any slot.
	// Removes Executioner if Town doesn't exist.
	if len(rl.mafia) == 0 {
		neutralEvil = removeUnique("Turncoat(Mafia)", neutralEvil)
	}
	if len(rl.coven) == 0 {
		neutralEvil = removeUnique("Turncoat(Coven)", neutralEvil)
	}
	if len(rl.coven) > 0 || (c.A > 0 && t.AnyCov) {
		neutralEvil = removeUnique("Witch", neutralEvil)
	}
	if c.TI+c.TP+c.TS+c.TK+c.RT == 0 {
		neutralEvil = removeUnique("Executioner", neutralEvil)
	}

	// Converts Vampire slots to Random Neutral if banned.
	if slices.Contains(ban, "vampire") && c.Vamp > 0 {
		fmt.Printf("Vampires banned, converting %v Vampire slots to Random Neutral.\n", c.Vamp)
		c.RN += c.Vamp
		c.Vamp = 0
	}
	// Adds guaranteed Vampires.
	if c.Vamp > 0 {
		fmt.Printf("Adding %v Vampires.\n", c.Vamp)
		for range c.Vamp {
			_, _, rl.neutral = insertGuaranteedRole(c.Vamp, neutralChaos, rl.neutral, "Vampire")
		}
	}

	// Converts Neutral subcategories to Random Neutral if no roles are available in the subcategory and skips role selection.
	if len(neutralKilling) == 0 && c.NK > 0 {
		fmt.Printf("No valid Neutral Killing roles, %v slots converted to Random Neutral.\n", c.NK)
		c.RN += c.NK
		c.NK = 0
	}
	if len(neutralChaos) == 0 && c.NC > 0 {
		fmt.Printf("No valid Neutral Chaos roles, %v slots converted to Random Neutral.\n", c.NC)
		c.RN += c.NC
		c.NC = 0
	}
	if len(neutralEvil) == 0 && c.NE > 0 {
		fmt.Printf("No valid Neutral Evil roles, %v slots converted to Random Neutral.\n", c.NE)
		c.RN += c.NE
		c.NE = 0
	}
	if len(neutralBenign) == 0 && c.NB > 0 {
		fmt.Printf("No valid Neutral Benign roles, %v slots converted to Random Neutral.\n", c.NB)
		c.RN += c.NB
		c.NB = 0
	}

	// Adds Neutral Killing and Neutral Chaos roles requested.
	if c.NK > 0 {
		fmt.Printf("Adding %v Neutral Killing.\n", c.NK)
		neutralKilling, rl.neutral, extra = randomRoleSelection(c.NK, neutralKilling, unique, rl.neutral)
		if extra > 0 {
			fmt.Printf("Converting %v slots to Random Neutral.\n", extra)
			c.RN += extra
		}
	}
	if c.NC > 0 {
		fmt.Printf("Adding %v Neutral Chaos.\n", c.NC)
		neutralChaos, rl.neutral, extra = randomRoleSelection(c.NC, neutralChaos, unique, rl.neutral)
		if extra > 0 {
			fmt.Printf("Converting %v slots to Random Neutral.\n", extra)
			c.RN += extra
		}
	}

	// Adds rolled Neutral roles to Guardian Angel target list.
	if len(rl.neutral) > 0 {
		rl.gaList = append(rl.gaList, rl.neutral...)
	}

	// Adds Neutral Evil, Neutral Benign, and Random Neutral roles requested, then adds eligible ones to Guardian Angel target list.
	numRoles := len(rl.neutral)
	if c.NE > 0 {
		fmt.Printf("Adding %v Neutral Evil.\n", c.NE)
		neutralEvil, rl.neutral, extra = randomRoleSelection(c.NE, neutralEvil, unique, rl.neutral)
		if extra > 0 {
			fmt.Printf("Converting %v slots to Random Neutral.\n", extra)
			c.RN += extra
		}
	}
	if c.NB > 0 {
		fmt.Printf("Adding %v Neutral Benign.\n", c.NB)
		neutralBenign, rl.neutral, extra = randomRoleSelection(c.NB, neutralBenign, unique, rl.neutral)
		if extra > 0 {
			fmt.Printf("Converting %v slots to Random Neutral.\n", extra)
			c.RN += extra
		}
	}

	randomNeutral := slices.Concat(neutralKilling, neutralChaos, neutralEvil, neutralBenign)
	if len(randomNeutral) == 0 && c.RN > 0 {
		fmt.Printf("No valid Neutral roles, %v slots converted to Any.\n", c.RN)
		c.A += c.RN
		c.RN = 0
	}
	if c.RN > 0 {
		fmt.Printf("Adding %v Random Neutral.\n", c.RN)
		randomNeutral, rl.neutral, extra = randomRoleSelection(c.RN, randomNeutral, unique, rl.neutral)
		if extra > 0 {
			fmt.Printf("Converting %v slots to Any.\n", extra)
			c.A += extra
		}
	}
	for i := numRoles; i < len(rl.neutral); i++ {
		if !slices.Contains(nonGA, rl.neutral[i]) {
			rl.gaList = append(rl.gaList, rl.neutral[i])
		}
	}

	// If Vampires exist, adds Vampire Hunter to the Town Killing list if it's not banned.
	if slices.Contains(rl.neutral, "Vampire") && !slices.Contains(ban, "vampire_hunter") {
		townKilling = append(townKilling, "Vampire_Hunter")
	}

	// Adds Jailor if guaranteed.
	if t.Jailor && c.TK > 0 {
		fmt.Println("Adding Jailor.")
		c.TK, townKilling, rl.town = insertGuaranteedRole(c.TK, townKilling, rl.town, "Jailor")
	} else if t.Jailor && c.RT > 0 {
		fmt.Println("Adding Jailor.")
		c.RT, townKilling, rl.town = insertGuaranteedRole(c.RT, townKilling, rl.town, "Jailor")
	}

	// Converts Town subcategories to Random Town if no roles are available in the subcategory and skips role selection.
	if len(townInvestigative) == 0 && c.TI > 0 {
		fmt.Printf("No valid Town Investigative roles, %v slots converted to Random Town.\n", c.TI)
		c.RT += c.TI
		c.TI = 0
	}
	if len(townProtective) == 0 && c.TP > 0 {
		fmt.Printf("No valid Town Protective roles, %v slots converted to Random Town.\n", c.TP)
		c.RT += c.TP
		c.TP = 0
	}
	if len(townSupport) == 0 && c.TS > 0 {
		fmt.Printf("No valid Town Support roles, %v slots converted to Random Town.\n", c.TS)
		c.RT += c.TS
		c.TS = 0
	}
	if len(townKilling) == 0 && c.TK > 0 {
		fmt.Printf("No valid Town Killing roles, %v slots converted to Random Town.\n", c.TK)
		c.RT += c.TK
		c.TK = 0
	}

	// Adds all Town roles requested.
	if c.TI > 0 {
		fmt.Printf("Adding %v Town Investigative.\n", c.TI)
		townInvestigative, rl.town, extra = randomRoleSelection(c.TI, townInvestigative, unique, rl.town)
		if extra > 0 {
			fmt.Printf("Converting %v slots to Random Town.\n", extra)
			c.RT += extra
		}
	}
	if c.TP > 0 {
		fmt.Printf("Adding %v Town Protective.\n", c.TP)
		townProtective, rl.town, extra = randomRoleSelection(c.TP, townProtective, unique, rl.town)
		if extra > 0 {
			fmt.Printf("Converting %v slots to Random Town.\n", extra)
			c.RT += extra
		}
	}
	if c.TS > 0 {
		fmt.Printf("Adding %v Town Support.\n", c.TS)
		townSupport, rl.town, extra = randomRoleSelection(c.TS, townSupport, unique, rl.town)
		if extra > 0 {
			fmt.Printf("Converting %v slots to Random Town.\n", extra)
			c.RT += extra
		}
	}
	if c.TK > 0 {
		fmt.Printf("Adding %v Town Killing.\n", c.TK)
		townKilling, rl.town, extra = randomRoleSelection(c.TK, townKilling, unique, rl.town)
		if extra > 0 {
			fmt.Printf("Converting %v slots to Random Town.\n", extra)
			c.RT += extra
		}
	}

	randomTown := slices.Concat(townInvestigative, townProtective, townSupport, townKilling)
	if len(randomTown) == 0 && c.RT > 0 {
		fmt.Printf("No valid Random Town roles, %v slots converted to Any.\n", c.A)
		c.A += c.RT
		c.RT = 0
	}
	if c.RT > 0 {
		fmt.Printf("Adding %v Random Town.\n", c.RT)
		randomTown, rl.town, extra = randomRoleSelection(c.RT, randomTown, unique, rl.town)
		if extra > 0 {
			fmt.Printf("Converting %v slots to Any.\n", extra)
			c.A += extra
		}
	}

	// Adds rolled Town roles to Guardian Angel target list.
	if len(rl.town) > 0 {
		rl.gaList = append(rl.gaList, rl.town...)
	}
	// Adds eligible Town roles to Executioner target list.
	for i := range rl.town {
		if !slices.Contains(nonExe, rl.town[i]) {
			rl.exeList = append(rl.exeList, rl.town[i])
		}
	}

	// Adds all Any roles requested.
	anyRole := slices.Concat(randomTown, randomNeutral)
	if t.AnyMaf {
		anyRole = slices.Concat(anyRole, randomMafia)
	}
	if t.AnyCov {
		anyRole = slices.Concat(anyRole, covenEvil)
	}
	if len(anyRole) == 0 && c.A > 0 {
		fmt.Printf("No valid Any roles, %v slots removed.\n", c.A)
		c.A = 0
	}
	if c.A > 0 {
		fmt.Printf("Adding %v Any.\n", c.A)
		_, rl.allAny = anyRoleSelection(c.A, anyRole, unique, randomTown, nonExe, randomMafia, covenEvil, ban, rl.allAny, t.Custom)
	}

	// Adds eligible Any roles to the GA and Executioner target lists
	for i := range rl.allAny {
		if !slices.Contains(nonExe, rl.allAny[i]) && (slices.Contains(randomTown, rl.allAny[i]) || slices.Contains(uniqueExe, rl.allAny[i])) {
			rl.exeList = append(rl.exeList, rl.allAny[i])
		}
		if !slices.Contains(nonGA, rl.allAny[i]) {
			rl.gaList = append(rl.gaList, rl.allAny[i])
		}
	}

	// Checks if Mafia only appeared in an Any slot and ensures that a Godfather or Mafioso exists.
	// Replaces the first Mafia on the list with either Godfather or Mafioso if one does not already exist.
	if len(rl.mafia) == 0 && t.AnyMaf && len(rl.allAny) > 0 && !slices.Contains(rl.allAny, "Godfather") && !slices.Contains(rl.allAny, "Mafioso") {
		for i := range rl.allAny {
			if slices.Contains(randomMafia, rl.allAny[i]) {
				if slices.Contains(ban, "godfather") {
					rl.allAny[i] = "Mafioso"
					break
				}
				if slices.Contains(ban, "mafioso") {
					rl.allAny[i] = "Godfather"
					break
				}
				randInt := rand.Intn(2)
				if randInt == 0 {
					rl.allAny[i] = "Godfather"
				} else {
					rl.allAny[i] = "Mafioso"
				}
				break
			}
		}
	}

	// If Executioner is rolled, checks if a valid target was also rolled.
	// If no valid targets, all Executioners are converted to Jesters.
	if slices.Contains(rl.neutral, "Executioner") && len(rl.exeList) == 0 {
		for i := range rl.neutral {
			if rl.neutral[i] == "Executioner" {
				rl.neutral[i] = "Jester"
			}
		}
		fmt.Println("No valid Executioner targets, converting to Jester")
	}
	if slices.Contains(rl.allAny, "Executioner") && len(rl.exeList) == 0 {
		for i := range rl.allAny {
			if rl.allAny[i] == "Executioner" {
				rl.allAny[i] = "Jester"
			}
		}
		fmt.Println("No valid Executioner targets, converting to Jester")
	}

	// If Guardian Angel is rolled, checks if a valid target was also rolled.
	// If no valid targets, all Guardian Angels are converted to Survivors.
	if slices.Contains(rl.neutral, "Guardian_Angel") && len(rl.gaList) == 0 {
		for i := range rl.neutral {
			if rl.neutral[i] == "Guardian_Angel" {
				rl.neutral[i] = "Survivor"
			}
		}
		fmt.Println("No valid GA targets, converting to Survivor")
	}
	if slices.Contains(rl.allAny, "Guardian_Angel") && len(rl.gaList) == 0 {
		for i := range rl.allAny {
			if rl.allAny[i] == "Guardian_Angel" {
				rl.allAny[i] = "Survivor"
			}
		}
		fmt.Println("No valid GA targets, converting to Survivor")
	}

	// Labels Executioner and Guardian Angel targets for roles appearing multiple times.
	if len(rl.exeList) > 0 {
		rl.exeList = labelTargets(rl.exeList)
	}
	if len(rl.gaList) > 0 {
		rl.gaList = labelTargets(rl.gaList)
	}
	// Assigns Executioner and Guardian Angel targets.
	if slices.Contains(rl.neutral, "Executioner") {
		rl.exeTargets = addTargets(rl.neutral, rl.exeList, rl.exeTargets, "Executioner")
	}
	if slices.Contains(rl.neutral, "Guardian_Angel") {
		rl.gaTargets = addTargets(rl.neutral, rl.gaList, rl.gaTargets, "Guardian_Angel")
	}
	if slices.Contains(rl.allAny, "Executioner") {
		rl.exeTargets = addTargets(rl.allAny, rl.exeList, rl.exeTargets, "Executioner")
	}
	if slices.Contains(rl.allAny, "Guardian_Angel") {
		rl.gaTargets = addTargets(rl.allAny, rl.gaList, rl.gaTargets, "Guardian_Angel")
	}

	return rl
}

// Removes Unique roles from the role category when the role is added to the list.
func removeUnique(role string, rolelist []string) []string {
	i := 0
	for idx, item := range rolelist {
		if item != role {
			rolelist[i] = rolelist[idx]
			i++
		}
	}
	return rolelist[:i]
}

// Randomly adds an eligible role to the role list and checks if it is unique.
func randomRoleSelection(num int, roleGroup, unique, roles []string) ([]string, []string, int) {
	for i := range num {
		if len(roleGroup) == 0 {
			fmt.Println("No valid roles left in category.")
			return roleGroup, roles, num - i
		}
		randomIdx := rand.Intn(len(roleGroup))
		randomRole := roleGroup[randomIdx]

		// Removes role from future rolls if Unique.
		if slices.Contains(unique, randomRole) {
			roleGroup = removeUnique(randomRole, roleGroup)
		}
		roles = append(roles, randomRole)
	}
	return roleGroup, roles, 0
}

// Randomly adds an any role to the role list, checks if unique, and checks if previously invalid roles are now valid options.
func anyRoleSelection(num int, roleGroup, unique, randomTown, nonExe, randomMafia, covenEvil, ban, roles []string, custom bool) ([]string, []string) {
	for i := range num {
		if len(roleGroup) == 0 {
			fmt.Printf("No valid roles left in category, %v slots removed.\n", num-i)
			return roleGroup, roles
		}
		randomIdx := rand.Intn(len(roleGroup))
		randomRole := roleGroup[randomIdx]

		// Adds Vampire Hunter to the role group if Vampire is rolled.
		if randomRole == "Vampire" && !slices.Contains(roleGroup, "Vampire_Hunter") && !slices.Contains(ban, "vampire_hunter") {
			roleGroup = append(roleGroup, "Vampire_Hunter")
		}

		// Adds Turncoat to role group if Mafia or Coven are rolled and custom roles are turned on.
		if slices.Contains(randomMafia, randomRole) && !slices.Contains(roleGroup, "Turncoat(Mafia)") && custom && !slices.Contains(ban, "turncoat(mafia)") {
			roleGroup = append(roleGroup, "Turncoat(Mafia)")
		}
		if slices.Contains(covenEvil, randomRole) && !slices.Contains(roleGroup, "Turncoat(Coven)") && custom && !slices.Contains(ban, "turncoat(coven)") {
			roleGroup = append(roleGroup, "Turncoat(Coven)")
		}

		// Adds Executioner to role group if eligible Town role is added.
		if slices.Contains(randomTown, randomRole) && !slices.Contains(roleGroup, "Executioner") && !slices.Contains(nonExe, randomRole) && !slices.Contains(ban, "executioner") {
			roleGroup = append(roleGroup, "Executioner")
		}

		// Removes role from future rolls if Unique.
		if slices.Contains(unique, randomRole) {
			roleGroup = removeUnique(randomRole, roleGroup)
		}

		roles = append(roles, randomRole)
	}
	return roleGroup, roles
}

// Adds a guaranteed role to the role list.
func insertGuaranteedRole(num int, roleGroup, roles []string, role string) (int, []string, []string) {
	roles = append(roles, role)
	roleGroup = removeUnique(role, roleGroup)
	num--
	return num, roleGroup, roles
}

// Removes banned roles from a role subcategory.
func checkBans(roleGroup, ban []string) []string {
	for i := len(roleGroup) - 1; i >= 0; i-- {
		if slices.Contains(ban, strings.ToLower(roleGroup[i])) {
			roleGroup = removeUnique(roleGroup[i], roleGroup)
		}
	}
	return roleGroup
}

// Assigns targets for Executioner and Guadian Angel roles.
func addTargets(roleGroup, eligibleList, targetList []string, role string) []string {
	for i := range roleGroup {
		if roleGroup[i] == role {
			randomIdx := rand.Intn(len(eligibleList))
			targetList = append(targetList, eligibleList[randomIdx])
		}
	}
	return targetList
}

// Distinguishes roles in target list if a role appears multiple times in the role list
func labelTargets(targetList []string) []string {
	for i := range targetList {
		role := targetList[i]
		if slices.Contains(targetList[i+1:], role) {
			targetList[i] = fmt.Sprintf("%v (1)", targetList[i])
			j := 2
			for k := i + 1; k < len(targetList); k++ {
				if targetList[k] == role {
					targetList[k] = fmt.Sprintf("%v (%v)", targetList[k], j)
					j++
				}
			}
		}
	}
	return targetList
}

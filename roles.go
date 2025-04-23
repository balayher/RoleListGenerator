package main

import (
	"math/rand"
	"slices"
)

func createRoles(ti, tp, ts, tk, rt, mk, ms, md, rm, ce, nk, nc, ne, nb, rn, a, vamp int, jailor, gf, cl, anyMaf, anyCov, anyVamp bool) ([]string, []string, []string, []string, []string) {
	// Initializes slices for each faction.
	town := []string{}
	mafia := []string{}
	coven := []string{}
	neutral := []string{}
	allAny := []string{}

	// Defines each Town role category.
	townInvestigative := []string{
		"Investigator",
		"Sheriff",
		"Lookout",
		"Tracker",
		"Psychic",
		"Spy",
		"Seer",
		"Detective",
	}
	townProtective := []string{
		"Bodyguard",
		"Doctor",
		"Crusader",
		"Trapper",
		"Cleric",
		"Oracle",
	}
	townSupport := []string{
		"Mayor",
		"Escort",
		"Retributionist",
		"Medium",
		"Transporter",
		"Monarch",
		"Governor",
		"Prosecutor",
		"Jack_of_All_Trades",
		"Timeshifter",
	}
	townKilling := []string{
		"Jailor",
		"Veteran",
		"Vigilante",
		"Gambler",
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
		"Poppet",
	}
	mafiaSupport := []string{
		"Consort",
		"Blackmailer",
		"Consigliere",
		"Watcher",
		"Angler",
		"Underboss",
		"Bouncer",
	}
	mafiaDeception := []string{
		"Disguiser",
		"Forger",
		"Framer",
		"Hypnotist",
		"Janitor",
		"Stager",
	}

	// Defines the Coven role category.
	covenEvil := []string{
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
	}

	// Defines each Netural role category.
	neutralKilling := []string{
		"Arsonist",
		"Juggernaut",
		"Serial_Killer",
		"Werewolf",
		"Mutator",
		"Horticulturist",
		"Shapeshifter",
		"Shroud",
		"Bombardier",
		"Gargoyle",
	}
	neutralEvil := []string{
		"Executioner",
		"Jester",
		"Witch",
		"Turncoat(Mafia)",
		"Turncoat(Coven)",
	}
	neutralChaos := []string{
		"Pirate",
		"Plaguebearer",
		"Inquisitor",
		"Anarchist",
		"Quack",
		"Stalker",
	}
	neutralBenign := []string{
		"Amnesiac",
		"Guardian_Angel",
		"Survivor",
	}

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

	if anyVamp {
		neutralChaos = append(neutralChaos, "Vampire")
	}

	// Adds Godfather if guaranteed, else adds either Godfather or Mafioso if Mafia exists.
	if gf && mk > 0 {
		mk, mafiaKilling, mafia = insertGuaranteedRole(mk, mafiaKilling, mafia, "Godfather")
	} else if gf && rm > 0 {
		rm, mafiaKilling, mafia = insertGuaranteedRole(rm, mafiaKilling, mafia, "Godfather")
	} else if mk > 0 {
		_, mafia = randomRoleSelection(1, gfMafioso, unique, mafia)
		mk--
		mafiaKilling = removeUnique(mafia[0], mafiaKilling)
	} else if rm > 0 {
		_, mafia = randomRoleSelection(1, gfMafioso, unique, mafia)
		rm--
		mafiaKilling = removeUnique(mafia[0], mafiaKilling)
	}

	// Adds all other Mafia roles requested.
	mafiaKilling, mafia = randomRoleSelection(mk, mafiaKilling, unique, mafia)
	mafiaDeception, mafia = randomRoleSelection(md, mafiaDeception, unique, mafia)
	mafiaSupport, mafia = randomRoleSelection(ms, mafiaSupport, unique, mafia)
	randomMafia := slices.Concat(mafiaKilling, mafiaDeception, mafiaSupport)
	randomMafia, mafia = randomRoleSelection(rm, randomMafia, unique, mafia)

	// Adds Coven Leader if guaranteed, then adds all other Coven roles requested.
	if cl && ce > 0 {
		ce, covenEvil, coven = insertGuaranteedRole(ce, covenEvil, coven, "Coven_Leader")
	}
	covenEvil, coven = randomRoleSelection(ce, covenEvil, unique, coven)

	// Removes Turncoats from the NE list if either Mafia or Coven doesn't exist.
	// Removes Witch from the NE list if Coven exists or can be rolled in an Any slot.
	if len(mafia) == 0 {
		neutralEvil = removeUnique("Turncoat(Mafia)", neutralEvil)
	}
	if len(coven) == 0 {
		neutralEvil = removeUnique("Turncoat(Coven)", neutralEvil)
	}
	if len(coven) > 0 || (a > 0 && anyCov) {
		neutralEvil = removeUnique("Witch", neutralEvil)
	}

	for i := 0; i < vamp; i++ {
		_, _, neutral = insertGuaranteedRole(vamp, neutralChaos, neutral, "Vampire")
	}

	// Adds all Neutral roles requested.
	neutralKilling, neutral = randomRoleSelection(nk, neutralKilling, unique, neutral)
	neutralChaos, neutral = randomRoleSelection(nc, neutralChaos, unique, neutral)
	neutralEvil, neutral = randomRoleSelection(ne, neutralEvil, unique, neutral)
	neutralBenign, neutral = randomRoleSelection(nb, neutralBenign, unique, neutral)
	randomNeutral := slices.Concat(neutralKilling, neutralChaos, neutralEvil, neutralBenign)
	randomNeutral, neutral = randomRoleSelection(rn, randomNeutral, unique, neutral)

	// If Vampires exist, adds Vampire Hunter to the Town Killing list.
	if slices.Contains(neutral, "Vampire") {
		townKilling = append(townKilling, "Vampire_Hunter")
	}

	// Adds Jailor if guaranteed.
	if jailor && tk > 0 {
		tk, townKilling, town = insertGuaranteedRole(tk, townKilling, town, "Jailor")
	} else if jailor && rt > 0 {
		rt, townKilling, town = insertGuaranteedRole(rt, townKilling, town, "Jailor")
	}

	// Adds all other Town roles requested.
	townInvestigative, town = randomRoleSelection(ti, townInvestigative, unique, town)
	townProtective, town = randomRoleSelection(tp, townProtective, unique, town)
	townSupport, town = randomRoleSelection(ts, townSupport, unique, town)
	townKilling, town = randomRoleSelection(tk, townKilling, unique, town)
	randomTown := slices.Concat(townInvestigative, townProtective, townSupport, townKilling)
	randomTown, town = randomRoleSelection(rt, randomTown, unique, town)

	// Adds all Any roles requested.
	anyRole := slices.Concat(randomTown, randomNeutral)
	if anyMaf {
		anyRole = slices.Concat(anyRole, randomMafia)
	}
	if anyCov {
		anyRole = slices.Concat(anyRole, covenEvil)
	}
	_, allAny = anyRoleSelection(a, anyRole, unique, randomMafia, covenEvil, allAny)

	// Checks if Mafia only appeared in an Any slot and ensures that a Godfather or Mafioso exists.
	// Replaces the first Mafia on the list with either Godfather or Mafioso if one does not already exist.
	if len(mafia) == 0 && anyMaf && len(allAny) > 0 && !slices.Contains(allAny, "Godfather") && !slices.Contains(allAny, "Mafia") {
		for i := 0; i < len(allAny); i++ {
			if slices.Contains(randomMafia, allAny[i]) {
				randInt := rand.Intn(2)
				if randInt == 0 {
					allAny[i] = "Godfather"
				} else {
					allAny[i] = "Mafioso"
				}
				break
			}
		}
	}

	return town, mafia, coven, neutral, allAny
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
func randomRoleSelection(num int, roleGroup, unique, roles []string) ([]string, []string) {
	for i := 0; i < num; i++ {
		randomIdx := rand.Intn(len(roleGroup))
		randomRole := roleGroup[randomIdx]
		// Removes role from future rolls if Unique.
		if slices.Contains(unique, randomRole) {
			roleGroup = removeUnique(randomRole, roleGroup)
		}
		roles = append(roles, randomRole)
	}
	return roleGroup, roles
}

// Randomly adds an any role to the role list, checks if unique, and checks if previously invalid roles are now valid options.
func anyRoleSelection(num int, roleGroup, unique, randomMafia, covenEvil, roles []string) ([]string, []string) {
	for i := 0; i < num; i++ {
		randomIdx := rand.Intn(len(roleGroup))
		randomRole := roleGroup[randomIdx]
		// Adds Vampire Hunter to the role group if Vampire is rolled.
		if randomRole == "Vampire" && !slices.Contains(roleGroup, "Vampire_Hunter") {
			roleGroup = append(roleGroup, "Vampire_Hunter")
		}
		// Adds Turncoat to role group if Mafia or Coven are rolled.
		if slices.Contains(randomMafia, randomRole) && !slices.Contains(roleGroup, "Turncoat(Mafia)") {
			roleGroup = append(roleGroup, "Turncoat(Mafia)")
		}
		if slices.Contains(covenEvil, randomRole) && !slices.Contains(roleGroup, "Turncoat(Coven)") {
			roleGroup = append(roleGroup, "Turncoat(Coven)")
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

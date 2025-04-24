package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

func createRoles(ti, tp, ts, tk, rt, mk, ms, md, rm, ce, nk, nc, ne, nb, rn, a, vamp int, jailor, gf, cl, anyMaf, anyCov, anyVamp, custom bool, ban []string) ([]string, []string, []string, []string, []string) {
	// Initializes slices for each faction for the final role list.
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
	covenEvil = checkBans(covenEvil, ban)

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
	if custom {
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

	// Adds Vampires to random pool if allowed
	if anyVamp && !slices.Contains(ban, "vampire") {
		neutralChaos = append(neutralChaos, "Vampire")
	}

	// Cancels guaranteed roles if they are on the ban list
	if slices.Contains(ban, "jailor") {
		jailor = false
	}
	if slices.Contains(ban, "godfather") {
		gf = false
	}
	if slices.Contains(ban, "coven_leader") {
		cl = false
	}

	// Checks if both Godfather and Mafioso are on the ban list and converts all Mafia slots to Any slots
	if len(gfMafioso) == 0 {
		maf := mk + ms + md + rm
		fmt.Printf("Both Godfather and Mafioso banned, converting %v mafia slots to Any.\n", maf)
		a += maf
		mk = 0
		ms = 0
		md = 0
		rm = 0
		anyMaf = false
	}

	// Adds Godfather if guaranteed, else adds either Godfather or Mafioso if Mafia exists.
	if gf && mk > 0 {
		fmt.Println("Adding Godfather.")
		mk, mafiaKilling, mafia = insertGuaranteedRole(mk, mafiaKilling, mafia, "Godfather")
	} else if gf && rm > 0 {
		fmt.Println("Adding Godfather.")
		rm, mafiaKilling, mafia = insertGuaranteedRole(rm, mafiaKilling, mafia, "Godfather")
	} else if mk > 0 {
		fmt.Println("Adding Godfather or Mafioso.")
		_, mafia = randomRoleSelection(1, gfMafioso, unique, mafia)
		mk--
		mafiaKilling = removeUnique(mafia[0], mafiaKilling)
	} else if rm > 0 {
		fmt.Println("Adding Godfather or Mafioso.")
		_, mafia = randomRoleSelection(1, gfMafioso, unique, mafia)
		rm--
		mafiaKilling = removeUnique(mafia[0], mafiaKilling)
	}

	// Converts Mafia subcategories to Random Mafia if no roles are available in the subcategory and skips role selection.
	if len(mafiaKilling) == 0 && mk > 0 {
		fmt.Printf("No valid Mafia Killing roles, %v slots converted to Random Mafia.\n", mk)
		rm += mk
		mk = 0
	}
	if len(mafiaDeception) == 0 && md > 0 {
		fmt.Printf("No valid Mafia Deception roles, %v slots converted to Random Mafia.\n", md)
		rm += md
		md = 0
	}
	if len(mafiaSupport) == 0 && ms > 0 {
		fmt.Printf("No valid Mafia Support roles, %v slots converted to Random Mafia.\n", ms)
		rm += ms
		ms = 0
	}

	// Adds all Mafia roles requested.
	if mk > 0 {
		fmt.Printf("Adding %v Mafia Killing.\n", mk)
		mafiaKilling, mafia = randomRoleSelection(mk, mafiaKilling, unique, mafia)
	}
	if md > 0 {
		fmt.Printf("Adding %v Mafia Deception.\n", md)
		mafiaDeception, mafia = randomRoleSelection(md, mafiaDeception, unique, mafia)
	}
	if ms > 0 {
		fmt.Printf("Adding %v Mafia Support.\n", ms)
		mafiaSupport, mafia = randomRoleSelection(ms, mafiaSupport, unique, mafia)
	}

	randomMafia := slices.Concat(mafiaKilling, mafiaDeception, mafiaSupport)
	if len(randomMafia) == 0 && rm > 0 {
		fmt.Printf("No valid Mafia roles, %v slots converted to Any.\n", rm)
		a += rm
		rm = 0
	}
	if rm > 0 {
		fmt.Printf("Adding %v Random Mafia.\n", rm)
		randomMafia, mafia = randomRoleSelection(rm, randomMafia, unique, mafia)
	}

	// Converts Coven slots to Any if all Coven roles are banned.
	if len(covenEvil) == 0 && ce > 0 {
		fmt.Printf("No valid Coven roles, %v slots converted to Any.\n", ce)
		a += ce
		ce = 0
	}

	// Adds Coven Leader if guaranteed, then adds all Coven roles requested.
	if cl && ce > 0 {
		fmt.Println("Adding Coven Leader.")
		ce, covenEvil, coven = insertGuaranteedRole(ce, covenEvil, coven, "Coven_Leader")
	}
	if ce > 0 {
		fmt.Printf("Adding %v Coven Evil.\n", ce)
		covenEvil, coven = randomRoleSelection(ce, covenEvil, unique, coven)
	}

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

	// Converts Vampire slots to Random Neutral if banned.
	if slices.Contains(ban, "vampire") && vamp > 0 {
		fmt.Printf("Vampires banned, converting %v Vampire slots to Random Neutral.\n", vamp)
		rn += vamp
		vamp = 0
	}
	// Adds guaranteed Vampires.
	if vamp > 0 {
		fmt.Printf("Adding %v Vampires.\n", vamp)
		for range vamp {
			_, _, neutral = insertGuaranteedRole(vamp, neutralChaos, neutral, "Vampire")
		}
	}

	// Converts Neutral subcategories to Random Neutral if no roles are available in the subcategory and skips role selection.
	if len(neutralKilling) == 0 && nk > 0 {
		fmt.Printf("No valid Neutral Killing roles, %v slots converted to Random Neutral.\n", nk)
		rn += nk
		nk = 0
	}
	if len(neutralChaos) == 0 && nc > 0 {
		fmt.Printf("No valid Neutral Chaos roles, %v slots converted to Random Neutral.\n", nc)
		rn += nc
		nc = 0
	}
	if len(neutralEvil) == 0 && ne > 0 {
		fmt.Printf("No valid Neutral Evil roles, %v slots converted to Random Neutral.\n", ne)
		rn += ne
		ne = 0
	}
	if len(neutralBenign) == 0 && nb > 0 {
		fmt.Printf("No valid Neutral Benign roles, %v slots converted to Random Neutral.\n", nb)
		rn += nb
		nb = 0
	}

	// Adds all Neutral roles requested.
	if nk > 0 {
		fmt.Printf("Adding %v Neutral Killing.\n", nk)
		neutralKilling, neutral = randomRoleSelection(nk, neutralKilling, unique, neutral)
	}
	if nc > 0 {
		fmt.Printf("Adding %v Neutral Chaos.\n", nc)
		neutralChaos, neutral = randomRoleSelection(nc, neutralChaos, unique, neutral)
	}
	if ne > 0 {
		fmt.Printf("Adding %v Neutral Evil.\n", ne)
		neutralEvil, neutral = randomRoleSelection(ne, neutralEvil, unique, neutral)
	}
	if nb > 0 {
		fmt.Printf("Adding %v Neutral Benign.\n", nb)
		neutralBenign, neutral = randomRoleSelection(nb, neutralBenign, unique, neutral)
	}

	randomNeutral := slices.Concat(neutralKilling, neutralChaos, neutralEvil, neutralBenign)
	if len(randomNeutral) == 0 && rn > 0 {
		fmt.Printf("No valid Neutral roles, %v slots converted to Any.\n", rn)
		a += rn
		rn = 0
	}
	if rn > 0 {
		fmt.Printf("Adding %v Random Neutral.\n", rn)
		randomNeutral, neutral = randomRoleSelection(rn, randomNeutral, unique, neutral)
	}

	// If Vampires exist, adds Vampire Hunter to the Town Killing list.
	if slices.Contains(neutral, "Vampire") {
		townKilling = append(townKilling, "Vampire_Hunter")
	}

	// Adds Jailor if guaranteed.
	if jailor && tk > 0 {
		fmt.Println("Adding Jailor.")
		tk, townKilling, town = insertGuaranteedRole(tk, townKilling, town, "Jailor")
	} else if jailor && rt > 0 {
		fmt.Println("Adding Jailor.")
		rt, townKilling, town = insertGuaranteedRole(rt, townKilling, town, "Jailor")
	}

	// Converts Town subcategories to Random Town if no roles are available in the subcategory and skips role selection.
	if len(townInvestigative) == 0 && ti > 0 {
		fmt.Printf("No valid Town Investigative roles, %v slots converted to Random Town.\n", ti)
		rt += ti
		ti = 0
	}
	if len(townProtective) == 0 && tp > 0 {
		fmt.Printf("No valid Town Protective roles, %v slots converted to Random Town.\n", tp)
		rt += tp
		tp = 0
	}
	if len(townSupport) == 0 && ts > 0 {
		fmt.Printf("No valid Town Support roles, %v slots converted to Random Town.\n", ts)
		rt += ts
		ts = 0
	}
	if len(townKilling) == 0 && tk > 0 {
		fmt.Printf("No valid Town Killing roles, %v slots converted to Random Town.\n", tk)
		rt += tk
		tk = 0
	}

	// Adds all Town roles requested.
	if ti > 0 {
		fmt.Printf("Adding %v Town Investigative.\n", ti)
		townInvestigative, town = randomRoleSelection(ti, townInvestigative, unique, town)
	}
	if tp > 0 {
		fmt.Printf("Adding %v Town Protective.\n", tp)
		townProtective, town = randomRoleSelection(tp, townProtective, unique, town)
	}
	if ts > 0 {
		fmt.Printf("Adding %v Town Support.\n", ts)
		townSupport, town = randomRoleSelection(ts, townSupport, unique, town)
	}
	if tk > 0 {
		fmt.Printf("Adding %v Town Killing.\n", tk)
		townKilling, town = randomRoleSelection(tk, townKilling, unique, town)
	}

	randomTown := slices.Concat(townInvestigative, townProtective, townSupport, townKilling)
	if len(randomTown) == 0 && rt > 0 {
		fmt.Printf("No valid Random Town roles, %v slots converted to Any.\n", a)
		a += rt
		rt = 0
	}
	if rt > 0 {
		fmt.Printf("Adding %v Random Town.\n", rt)
		randomTown, town = randomRoleSelection(rt, randomTown, unique, town)
	}

	// Adds all Any roles requested.
	anyRole := slices.Concat(randomTown, randomNeutral)
	if anyMaf {
		anyRole = slices.Concat(anyRole, randomMafia)
	}
	if anyCov {
		anyRole = slices.Concat(anyRole, covenEvil)
	}
	if len(anyRole) == 0 && a > 0 {
		fmt.Printf("No valid Any roles, %v slots removed.\n", a)
		a = 0
	}
	if a > 0 {
		fmt.Printf("Adding %v Any.\n", a)
		_, allAny = anyRoleSelection(a, anyRole, unique, randomMafia, covenEvil, allAny, custom)
	}

	// Checks if Mafia only appeared in an Any slot and ensures that a Godfather or Mafioso exists.
	// Replaces the first Mafia on the list with either Godfather or Mafioso if one does not already exist.
	if len(mafia) == 0 && anyMaf && len(allAny) > 0 && !slices.Contains(allAny, "Godfather") && !slices.Contains(allAny, "Mafia") {
		for i := range allAny {
			if slices.Contains(randomMafia, allAny[i]) {
				if slices.Contains(ban, "godfather") {
					allAny[i] = "Mafioso"
					break
				}
				if slices.Contains(ban, "mafioso") {
					allAny[i] = "Godfather"
					break
				}
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
	for i := range num {
		if len(roleGroup) == 0 {
			fmt.Printf("No valid roles left in category, %v slots removed.\n", num-i)
			return roleGroup, roles
		}
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
func anyRoleSelection(num int, roleGroup, unique, randomMafia, covenEvil, roles []string, custom bool) ([]string, []string) {
	for i := range num {
		if len(roleGroup) == 0 {
			fmt.Printf("No valid roles left in category, %v slots removed.\n", num-i)
			return roleGroup, roles
		}
		randomIdx := rand.Intn(len(roleGroup))
		randomRole := roleGroup[randomIdx]

		// Adds Vampire Hunter to the role group if Vampire is rolled.
		if randomRole == "Vampire" && !slices.Contains(roleGroup, "Vampire_Hunter") {
			roleGroup = append(roleGroup, "Vampire_Hunter")
		}

		// Adds Turncoat to role group if Mafia or Coven are rolled and custom roles are turned on.
		if slices.Contains(randomMafia, randomRole) && !slices.Contains(roleGroup, "Turncoat(Mafia)") && custom {
			roleGroup = append(roleGroup, "Turncoat(Mafia)")
		}
		if slices.Contains(covenEvil, randomRole) && !slices.Contains(roleGroup, "Turncoat(Coven)") && custom {
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

// Removes banned roles from a role subcategory.
func checkBans(roleGroup, ban []string) []string {
	for i := len(roleGroup) - 1; i >= 0; i-- {
		if slices.Contains(ban, strings.ToLower(roleGroup[i])) {
			roleGroup = removeUnique(roleGroup[i], roleGroup)
		}
	}
	return roleGroup
}

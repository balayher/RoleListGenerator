package main

import (
	"math/rand"
	"slices"
)

func createRoles(ti, tp, ts, tk, rt, mk, ms, md, rm, ce int, jailor, gf, cl bool) ([]string, []string, []string) {
	town := []string{}
	mafia := []string{}
	coven := []string{}

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
		"Vampire_Hunter",
		"Gambler",
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

	unique := []string{
		"Cleric",
		"Oracle",
		"Mayor",
		"Retributionist",
		"Governor",
		"Prosecutor",
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
	}

	if gf && mk > 0 {
		gf, mk, mafiaKilling, mafia = insertGuaranteedRole(mk, mafiaKilling, mafia, "Godfather")
	}
	mafiaKilling, mafia = randomRoleSelection(mk, mafiaKilling, unique, mafia)
	mafiaDeception, mafia = randomRoleSelection(md, mafiaDeception, unique, mafia)
	mafiaSupport, mafia = randomRoleSelection(ms, mafiaSupport, unique, mafia)
	randomMafia := slices.Concat(mafiaKilling, mafiaDeception, mafiaSupport)
	if gf && rm > 0 && !slices.Contains(mafia, "Godfather") {
		gf, rm, randomMafia, mafia = insertGuaranteedRole(rm, randomMafia, mafia, "Godfather")
	}
	randomMafia, mafia = randomRoleSelection(rm, randomMafia, unique, mafia)

	if cl && ce > 0 {
		cl, ce, covenEvil, coven = insertGuaranteedRole(ce, covenEvil, coven, "Coven_Leader")
	}
	covenEvil, coven = randomRoleSelection(ce, covenEvil, unique, coven)

	if jailor && tk > 0 {
		jailor, tk, townKilling, town = insertGuaranteedRole(tk, townKilling, town, "Jailor")
	}
	townInvestigative, town = randomRoleSelection(ti, townInvestigative, unique, town)
	townProtective, town = randomRoleSelection(tp, townProtective, unique, town)
	townSupport, town = randomRoleSelection(ts, townSupport, unique, town)
	townKilling, town = randomRoleSelection(tk, townKilling, unique, town)
	randomTown := slices.Concat(townInvestigative, townProtective, townSupport, townKilling)
	if jailor && rt > 0 && !slices.Contains(town, "Jailor") {
		jailor, rt, randomTown, town = insertGuaranteedRole(rt, randomTown, town, "Jailor")
	}
	randomTown, town = randomRoleSelection(rt, randomTown, unique, town)

	return town, mafia, coven
}

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

func randomRoleSelection(num int, roleGroup, unique, roles []string) ([]string, []string) {
	for i := 0; i < num; i++ {
		randomIdx := rand.Intn(len(roleGroup))
		randomRole := roleGroup[randomIdx]
		if slices.Contains(unique, randomRole) {
			roleGroup = removeUnique(randomRole, roleGroup)
		}
		roles = append(roles, randomRole)
	}
	return roleGroup, roles
}

func insertGuaranteedRole(num int, roleGroup, roles []string, role string) (bool, int, []string, []string) {
	roles = append(roles, role)
	roleGroup = removeUnique(role, roleGroup)
	num--
	return false, num, roleGroup, roles
}

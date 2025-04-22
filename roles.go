package main

import (
	"math/rand"
	"slices"
)

func createRoles(ti, tp, ts, tk, rt, mk, ms, md, rm, ce, nk, nc, ne, nb, rn, a int, jailor, gf, cl bool) ([]string, []string, []string, []string, []string) {
	town := []string{}
	mafia := []string{}
	coven := []string{}
	neutral := []string{}
	allAny := []string{}

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
		"Vampire",
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

	if gf && mk > 0 {
		gf, mk, mafiaKilling, mafia = insertGuaranteedRole(mk, mafiaKilling, mafia, "Godfather")
	}
	mafiaKilling, mafia = randomRoleSelection(mk, mafiaKilling, unique, mafia)
	mafiaDeception, mafia = randomRoleSelection(md, mafiaDeception, unique, mafia)
	mafiaSupport, mafia = randomRoleSelection(ms, mafiaSupport, unique, mafia)
	randomMafia := slices.Concat(mafiaKilling, mafiaDeception, mafiaSupport)
	if gf && rm > 0 && !slices.Contains(mafia, "Godfather") {
		_, rm, randomMafia, mafia = insertGuaranteedRole(rm, randomMafia, mafia, "Godfather")
	}
	randomMafia, mafia = randomRoleSelection(rm, randomMafia, unique, mafia)

	if cl && ce > 0 {
		_, ce, covenEvil, coven = insertGuaranteedRole(ce, covenEvil, coven, "Coven_Leader")
	}
	covenEvil, coven = randomRoleSelection(ce, covenEvil, unique, coven)

	if len(mafia) == 0 {
		neutralEvil = removeUnique("Turncoat(Mafia)", neutralEvil)
	}
	if len(coven) == 0 {
		neutralEvil = removeUnique("Turncoat(Coven)", neutralEvil)
	}

	neutralKilling, neutral = randomRoleSelection(nk, neutralKilling, unique, neutral)
	neutralChaos, neutral = randomRoleSelection(nc, neutralChaos, unique, neutral)
	neutralEvil, neutral = randomRoleSelection(ne, neutralEvil, unique, neutral)
	neutralBenign, neutral = randomRoleSelection(nb, neutralBenign, unique, neutral)
	randomNeutral := slices.Concat(neutralKilling, neutralChaos, neutralEvil, neutralBenign)
	randomNeutral, neutral = randomRoleSelection(rn, randomNeutral, unique, neutral)

	if slices.Contains(neutral, "Vampire") {
		townKilling = append(townKilling, "Vampire_Hunter")
	}

	if jailor && tk > 0 {
		jailor, tk, townKilling, town = insertGuaranteedRole(tk, townKilling, town, "Jailor")
	}
	townInvestigative, town = randomRoleSelection(ti, townInvestigative, unique, town)
	townProtective, town = randomRoleSelection(tp, townProtective, unique, town)
	townSupport, town = randomRoleSelection(ts, townSupport, unique, town)
	townKilling, town = randomRoleSelection(tk, townKilling, unique, town)
	randomTown := slices.Concat(townInvestigative, townProtective, townSupport, townKilling)
	if jailor && rt > 0 && !slices.Contains(town, "Jailor") {
		_, rt, randomTown, town = insertGuaranteedRole(rt, randomTown, town, "Jailor")
	}
	randomTown, town = randomRoleSelection(rt, randomTown, unique, town)

	anyRole := slices.Concat(randomTown, randomMafia, randomNeutral, covenEvil)
	_, allAny = randomRoleSelection(a, anyRole, unique, allAny)

	return town, mafia, coven, neutral, allAny
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

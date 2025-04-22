package main

import (
	"math/rand"
	"slices"
)

func createRoles(ti, tp, ts, tk, rt int) []string {
	town := []string{}

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

	unique := []string{
		"Cleric",
		"Oracle",
		"Mayor",
		"Retributionist",
		"Governor",
		"Prosecutor",
		"Jailor",
		"Veteran",
	}

	for i := 0; i < ti; i++ {
		randomIdx := rand.Intn(len(townInvestigative))
		randomRole := townInvestigative[randomIdx]
		if slices.Contains(unique, randomRole) {
			townInvestigative = removeUnique(randomRole, townInvestigative)
		}
		town = append(town, randomRole)
	}

	for i := 0; i < tp; i++ {
		randomIdx := rand.Intn(len(townProtective))
		randomRole := townProtective[randomIdx]
		if slices.Contains(unique, randomRole) {
			townProtective = removeUnique(randomRole, townProtective)
		}
		town = append(town, randomRole)
	}

	for i := 0; i < ts; i++ {
		randomIdx := rand.Intn(len(townSupport))
		randomRole := townSupport[randomIdx]
		if slices.Contains(unique, randomRole) {
			townSupport = removeUnique(randomRole, townSupport)
		}
		town = append(town, randomRole)
	}

	for i := 0; i < tk; i++ {
		randomIdx := rand.Intn(len(townKilling))
		randomRole := townKilling[randomIdx]
		if slices.Contains(unique, randomRole) {
			townKilling = removeUnique(randomRole, townKilling)
		}
		town = append(town, randomRole)
	}

	randomTown := slices.Concat(townInvestigative, townProtective, townSupport, townKilling)

	for i := 0; i < rt; i++ {
		randomIdx := rand.Intn(len(randomTown))
		randomRole := randomTown[randomIdx]
		if slices.Contains(unique, randomRole) {
			randomTown = removeUnique(randomRole, randomTown)
		}
		town = append(town, randomRole)
	}

	return town
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

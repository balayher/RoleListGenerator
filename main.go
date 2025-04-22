package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var ti, tp, ts, tk, rt, mk, ms, md, rm, ce int
	jailor, gf, cl := false, false, false

	rand.New(rand.NewSource(time.Now().UnixNano()))

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
	if tk > 0 {
		fmt.Print("Do you want a guaranteed Jailor? ")
		jailor = getYesNo()
	}

	fmt.Print("Enter the number of Random Town: ")
	rt, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, RT set to 0")
	}
	if (tk == 0) && (rt > 0) {
		fmt.Print("Do you want a guaranteed Jailor? ")
		jailor = getYesNo()
	}

	fmt.Print("Enter the number of Mafia Killing: ")
	mk, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, MK set to 0")
	}
	if mk > 5 {
		fmt.Println("Maximum value exceeded, set to 5")
		mk = 5
	}
	if mk > 0 {
		fmt.Print("Do you want a guaranteed Godfather? ")
		gf = getYesNo()
	}

	fmt.Print("Enter the number of Mafis Support: ")
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
	if (mk == 0) && (rm > 0) {
		fmt.Print("Do you want a guaranteed Godfather? ")
		gf = getYesNo()
	}

	fmt.Print("Enter the number of Coven Evil: ")
	ce, err = getInput()
	if err != nil {
		fmt.Println("Invalid input, CE set to 0")
	}
	if ce > 10 {
		fmt.Println("Maximum value exceeded, set to 10")
		ce = 10
	}
	if ce > 0 {
		fmt.Print("Do you want a guaranteed Coven Leader? ")
		cl = getYesNo()
	}

	fmt.Printf("Your Town has %v TI, %v TP, %v TS, %v TK, and %v RT\n", ti, tp, ts, tk, rt)
	fmt.Printf("Your Mafia has %v MK, %v MS, %v MD, and %v RM\n", mk, ms, md, rm)

	town, mafia, coven := createRoles(ti, tp, ts, tk, rt, mk, ms, md, rm, ce, jailor, gf, cl)
	fmt.Println("Town:")
	fmt.Println(town)
	fmt.Println("Mafia:")
	fmt.Println(mafia)
	fmt.Println("Coven:")
	fmt.Println(coven)
}

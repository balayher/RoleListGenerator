package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var ti, tp, ts, tk, rt int

	rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Print("Enter the number of Town Investigative: ")
	fmt.Scan(&ti)
	fmt.Print("Enter the number of Town Protective: ")
	fmt.Scan(&tp)
	fmt.Print("Enter the number of Town Support: ")
	fmt.Scan(&ts)
	fmt.Print("Enter the number of Town Killing: ")
	fmt.Scan(&tk)
	fmt.Print("Enter the number of Random Town: ")
	fmt.Scan(&rt)
	fmt.Printf("You have %v TI, %v TP, %v TS, %v TK, and %v RT\n", ti, tp, ts, tk, rt)

	town := createRoles(ti, tp, ts, tk, rt)
	fmt.Println(town)
}

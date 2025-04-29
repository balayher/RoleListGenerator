package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func getJsonCounts(c Counts, filename string) (Counts, bool) {

	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Printf("Error opening %v\n", filename)
		return c, false
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("Error parsing data from %v\n", filename)
		return c, false
	}
	var jsonCounts Counts
	err = json.Unmarshal(data, &jsonCounts)
	if err != nil {
		fmt.Printf("Error parsing data from %v\n", filename)
		return c, false
	}

	return jsonCounts, true
}

func getJsonOptions(t Options, filename string) (Options, bool) {

	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Printf("Error opening %v\n", filename)
		return t, false
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("Error parsing data from %v\n", filename)
		return t, false
	}
	var jsonOptions Options
	err = json.Unmarshal(data, &jsonOptions)
	if err != nil {
		fmt.Printf("Error parsing data from %v\n", filename)
		return t, false
	}

	return jsonOptions, true
}

func saveJson(v any, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	b, err := json.MarshalIndent(&v, "", "\t")
	if err == nil {
		fmt.Fprint(f, string(b))
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v written successfully\n", filename)
}

func PrettyPrint(v any) (err error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err == nil {
		fmt.Println(string(b))
	}
	return err
}

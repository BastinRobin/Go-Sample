package main

import "os"

func main() {

	// Open a file
	file, err := os.Open("sample.csv")
	if err != nil {
		panic(err)
	}
	// Read it
	// Print it

}

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/onegkg/literator/internal/literator"
)

func main() {
	// Transliterate
	path := filepath.Join("test.txt")
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	processed := literator.Preprocess(string(dat))
	irOne := literator.ConvertToIROne(processed)
	irTwo, err := literator.ConvertToIRTwo(irOne)
	if err != nil {
		fmt.Printf("Error converting to IR2: %v\n", err)
		os.Exit(1)
	}
	// fmt.Printf("Input:\n%v\n", string(dat))
	// fmt.Printf("Processed:\n%v\n", processed)
	// fmt.Printf("IROne:\n%v\n", irOne.StringFromHead())
	// fmt.Printf("IRTwo:\n%v\n", irTwo.StringFromHead())

	literation := literator.Literate(irTwo)
	fmt.Printf("Transliteration:\n%v\n", literation)

}

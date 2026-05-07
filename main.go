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

	literation := literator.Literate(irTwo)
	fmt.Printf("%v\n", literation)

}

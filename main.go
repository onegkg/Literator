package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/onegkg/literator/internal/literator"
)

func main() {
	// get input
	// preprocess
	// Convert to IR
	// Transliterate
	path := filepath.Join("test.txt")
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error: %w\n", err)
		os.Exit(1)
	}
	processed := literator.Preprocess(string(dat))
	irOne := literator.ConvertToIROne(processed)
	irString := irOne.StringFromHead()
	fmt.Printf("%v\n", irOne.DebugPrint())
	fmt.Printf("%s", irString)

	fmt.Printf("\n")
}

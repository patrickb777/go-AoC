package main

import (
	"advent-of-code/readfile"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	// Init
	start := time.Now()
	fmt.Println("")
	f := flag.String("f", "none", "Input file")
	flag.Parse()

	// Parse Input
	inputFile := readfile.ReadFile(*f)
	fmt.Sprintln(inputFile) //Does not output to stdout

	// Processing

	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func parseDirections(input readfile.InputFile) {
}
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	context := initFromArgs()
	err := runScenario(context)
	if err != nil {
		err := fmt.Errorf("failed to run scenario '%s': %w", context.scenarioName, err)
		log.Fatal(err)
	}
}

func initFromArgs() *contextData {
	// read args
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		printUsage()
		os.Exit(0)
	}

	// parse args
	context, err := parseArgs(argsWithoutProg)
	if err != nil {
		err := fmt.Errorf("failed to parse arguments: %w", err)
		log.Fatal(err)
	}

	return context
}

func printUsage() {
	fmt.Println("Request")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("    rq <scenarioName> [[key=value]] [-e <envName>] [-v]")
}

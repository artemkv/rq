package main

import (
	"fmt"
	"strings"
)

func parseArgs(argsWithoutProg []string) (string, string, map[string]string, bool, error) {
	scenarioName := ""
	if len(argsWithoutProg) > 0 {
		scenarioName = argsWithoutProg[0]
	}

	env := ""
	inputs := make(map[string]string)
	isVerbose := false
	skip := false
	if len(argsWithoutProg) > 1 {
		for idx, arg := range argsWithoutProg[1:] {
			if skip {
				skip = false
				continue
			}

			// TODO: refactor a bit
			if arg == "-e" {
				if len(argsWithoutProg[1:]) <= idx+1 {
					return "", "", nil, false, fmt.Errorf("found '-e' but missing environment name")
				}
				env = argsWithoutProg[1:][idx+1]
				skip = true
			} else if arg == "-v" {
				isVerbose = true
			} else {
				parts := strings.Split(arg, "=")
				if len(parts) != 2 {
					return "", "", nil, false, fmt.Errorf("argument format '%s' is not recognized", arg)
				}
				inputs[parts[0]] = parts[1]
			}
		}
	}

	return scenarioName, env, inputs, isVerbose, nil
}

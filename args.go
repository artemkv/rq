package main

import (
	"fmt"
	"strings"
)

func parseArgs(argsWithoutProg []string) (*contextData, error) {
	context := &contextData{
		scenarioName: "",
		envName:      "",
		overrides:    make(map[string]string),
		isVerbose:    false,
	}

	if len(argsWithoutProg) > 0 {
		context.scenarioName = argsWithoutProg[0]
	}

	skip := false
	if len(argsWithoutProg) > 1 {
		restOfArgs := argsWithoutProg[1:]
		for idx, arg := range restOfArgs {
			if skip {
				skip = false
				continue
			}

			if arg == "-e" {
				if len(restOfArgs) <= idx+1 {
					return nil, fmt.Errorf("found '-e' but missing environment name")
				}
				context.envName = restOfArgs[idx+1]
				skip = true
			} else if arg == "-v" {
				context.isVerbose = true
			} else {
				parts := strings.Split(arg, "=")
				if len(parts) != 2 {
					return nil, fmt.Errorf("argument format '%s' is not recognized", arg)
				}
				context.overrides[parts[0]] = parts[1]
			}
		}
	}

	return context, nil
}

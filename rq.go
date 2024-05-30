package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// read args
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		printUsage()
		os.Exit(1)
	}

	// parse args
	scenarioName, envName, overrides, isVerbose, err := parseArgs(argsWithoutProg)
	if err != nil {
		err := fmt.Errorf("failed to parse arguments %w", err)
		log.Fatal(err)
	}
	if scenarioName == "" {
		printUsage()
		os.Exit(1)
	}

	// initialize environment
	inputs, err := loadEnvDefaults(envName)
	if err != nil {
		err := fmt.Errorf("failed to load environment %w", err)
		log.Fatal(err)
	}

	// merge inputs into the environment
	for k, v := range overrides {
		inputs[k] = v
	}

	// read config
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	// find the scenario
	scenario, ok := config.Scenarios[scenarioName]
	if !ok {
		err := fmt.Sprintf("scenario '%s' was not found.", scenarioName)
		log.Fatal(err)
	}

	// go through requests
	var finalBody []byte
	var finalResp *http.Response
	for _, requestId := range scenario.Sequence {
		// find the request
		request, ok := config.Requests[requestId]
		if !ok {
			err := fmt.Sprintf("scenario '%s' is referencing unknown request '%s'.", scenarioName, requestId)
			log.Fatal(err)
		}

		// parametrize the request
		parametrized := parametrize(request, inputs)

		// validate all the params were substituted
		// TODO:

		// verbose output
		if isVerbose {
			fmt.Println("==========================")
			fmt.Printf("Request: %s\n", requestId)
			fmt.Println("==========================")
			printRequest(parametrized)
			fmt.Println("--------------------------")
		}

		// make the request
		resp, err := makeRequest(parametrized)
		if err != nil {
			err := fmt.Errorf("failed to execute '%s': %w", requestId, err)
			log.Fatal(err)
		}

		// read body
		body, err := getBody(resp)
		if err != nil {
			err := fmt.Errorf("failed to read response of '%s': %w", requestId, err)
			log.Fatal(err)
		}

		// save for the final output
		finalBody = body
		finalResp = resp

		// process outputs
		if len(request.Outputs) > 0 {
			// extract outputs
			outputs, err := extractOutputs(body, request.Outputs)
			if err != nil {
				err := fmt.Errorf("failed to extract outputs of '%s': %w", requestId, err)
				log.Fatal(err)
			}

			// feed into the next request inputs
			for k, v := range outputs {
				inputs[k] = v
			}
		}

		// verbose output
		if isVerbose {
			err = printResponse(resp, body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("==========================")
		}
	}

	// print the final output
	if !isVerbose {
		err = printResponse(finalResp, finalBody)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func printUsage() {
	// TODO:
	fmt.Println("Request")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("    rq <scenarioName>")
}

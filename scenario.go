package main

import (
	"fmt"
	"net/http"
)

type contextData struct {
	scenarioName string
	envName      string
	overrides    map[string]string
	isVerbose    bool
}

func runScenario(context *contextData) error {
	// read environment defaults
	inputs, err := readEnvDefaults(context.envName)
	if err != nil {
		return fmt.Errorf("failed to load environment defaults: %w", err)
	}

	// read scenarios
	scenarioDescription, err := readScenarioDescription()
	if err != nil {
		return fmt.Errorf("failed to read scenario file: %w", err)
	}

	// merge overrides into the defaults
	for k, v := range context.overrides {
		inputs[k] = v
	}

	// find the scenario
	scenario, ok := scenarioDescription.Scenarios[context.scenarioName]
	if !ok {
		return fmt.Errorf("scenario '%s' was not found.", context.scenarioName)
	}

	// go through requests
	var finalBody []byte
	var finalResp *http.Response
	for _, requestId := range scenario.Sequence {
		// find the request
		request, ok := scenarioDescription.Requests[requestId]
		if !ok {
			return fmt.Errorf("scenario '%s' is referencing unknown request '%s'.",
				context.scenarioName, requestId)
		}

		// parametrize the request
		parametrized := parametrize(request, inputs)

		// validate all the params were substituted
		// TODO:

		// verbose output
		if context.isVerbose {
			verbosePrintRequest(requestId, parametrized)
		}

		// make the request
		resp, err := makeRequest(parametrized)
		if err != nil {
			return fmt.Errorf("failed to execute '%s': %w", requestId, err)
		}

		// read body
		body, err := readBody(resp)
		if err != nil {
			return fmt.Errorf("failed to read response of '%s': %w", requestId, err)
		}

		// save for the final output
		finalBody = body
		finalResp = resp

		// process outputs
		if len(request.Outputs) > 0 {
			// extract outputs
			outputs, err := extractOutputs(body, request.Outputs)
			if err != nil {
				return fmt.Errorf("failed to extract outputs of '%s': %w", requestId, err)
			}

			// feed into the next request inputs
			for k, v := range outputs {
				inputs[k] = v
			}
		}

		// verbose output
		if context.isVerbose {
			err = verbosePrintResponse(resp, body)
			if err != nil {
				return fmt.Errorf("failed to print response: %w", err)
			}
		}
	}

	// print the final output
	if !context.isVerbose {
		err := printResponse(finalResp, finalBody)
		if err != nil {
			return fmt.Errorf("failed to print response: %w", err)
		}
	}

	return nil
}

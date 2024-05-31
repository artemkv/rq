package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type scenarioDescriptionData struct {
	Scenarios map[string]scenarioData `json:"scenarios"`
	Requests  map[string]requestData  `json:"requests"`
}

type scenarioData struct {
	Sequence []string `json:"seq"`
}

type requestData struct {
	Method  string              `json:"method"`
	Url     string              `json:"url"`
	Headers map[string]string   `json:"headers"`
	Body    string              `json:"body"`
	Outputs []requestOutputData `json:"output"`
}

type requestOutputData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func readScenarioDescription() (*scenarioDescriptionData, error) {
	scenarioJson, err := os.ReadFile("rq.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read 'rq.json': %w", err)
	}

	var scenarioDescription scenarioDescriptionData
	err = json.Unmarshal(scenarioJson, &scenarioDescription)
	if err != nil {
		return nil, fmt.Errorf("failed to parse 'rq.json': %w", err)
	}

	return &scenarioDescription, nil
}

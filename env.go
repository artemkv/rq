package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type envData struct {
	Inputs []envInputData `json:"inputs"`
}

type envInputData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func loadEnvDefaults(name string) (map[string]string, error) {
	if name == "" {
		return make(map[string]string), nil
	}

	envFileName := fmt.Sprintf("%s.env.json", name)

	envJson, err := os.ReadFile(envFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read '%s': %w", envFileName, err)
	}

	var env envData
	err = json.Unmarshal(envJson, &env)
	if err != nil {
		return nil, fmt.Errorf("failed to parse '%s': %w", envFileName, err)
	}

	envDefaults := make(map[string]string, len(env.Inputs))
	for _, input := range env.Inputs {
		envDefaults[input.Key] = input.Value
	}

	return envDefaults, nil
}

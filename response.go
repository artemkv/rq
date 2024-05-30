package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/PaesslerAG/jsonpath"
)

func getBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func printResponse(resp *http.Response, body []byte) error {
	fmt.Printf("%v\n", resp.Status)
	for header := range resp.Header {
		for _, headerValue := range resp.Header[header] {
			fmt.Printf("%s: %s\n", header, headerValue)
		}
	}
	fmt.Println("")
	fmt.Printf("%s\n", string(body[:]))
	return nil
}

func extractOutputs(body []byte, outputs []requestOutputData) (map[string]string, error) {
	bodyJson := interface{}(nil)

	err := json.Unmarshal(body, &bodyJson)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body as JSON: %w", err)
	}

	extracted := make(map[string]string)
	for _, output := range outputs {
		val, err := jsonpath.Get(output.Value, bodyJson)
		if err != nil {
			return nil, fmt.Errorf("failed to extract output '%s': %w", output.Value, err)
		}
		extracted[output.Key] = fmt.Sprint(val)
	}

	return extracted, nil
}

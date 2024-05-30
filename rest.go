package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func Get(url string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	for header := range headers {
		req.Header.Add(header, headers[header])
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return resp, nil
}

func Post(url string, headers map[string]string, data []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	for header := range headers {
		req.Header.Add(header, headers[header])
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return resp, nil
}

func Put(url string, headers map[string]string, data []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	for header := range headers {
		req.Header.Add(header, headers[header])
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return resp, nil
}

func Delete(url string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	for header := range headers {
		req.Header.Add(header, headers[header])
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return resp, nil
}
